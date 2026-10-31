package handlers

import (
	"context"
	"encoding/json"
	"hs-ascii-live/engine"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type animationResponse struct {
	Name string `json:"name"`
	Category string `json:"category"`
	Description string `json:"description"`
}

func muxVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func AnimationHandler(w http.ResponseWriter, r *http.Request) {
	vars := muxVars(r)
	name := vars["name"]
	if name == "" {
		http.NotFound(w, r)
		return
	}
	anim, ok := engine.GlobalRegistry.Get(name)
	if !ok {
		http.Error(w, "Animation not found", http.StatusNotFound)
		return
	}
	opts := parseOptions(r)
	// Set headers for streaming
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Connection", "keep-alive")

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Detect client disconnect
	go func() {
		<-r.Context().Done()
		cancel()
	}()

	// Stream
	err := engine.StreamAnimation(ctx, w, anim, opts)
	if err != nil && err != context.Canceled {
		// log but don't panic
	}
}

func parseOptions(r *http.Request) engine.AnimationOptions {
	q := r.URL.Query()
	opts := engine.AnimationOptions{
		Speed: "normal",
		FPS:   10,
	}
	if s := q.Get("speed"); s != "" {
		opts.Speed = s
	}
	if fpsStr := q.Get("fps"); fpsStr != "" {
		if fps, err := strconv.Atoi(fpsStr); err == nil {
			opts.FPS = fps
		}
	}
	if color := q.Get("color"); color != "" {
		opts.Color = color
	}
	if wStr := q.Get("width"); wStr != "" {
		if w, err := strconv.Atoi(wStr); err == nil {
			opts.Width = w
		}
	}
	if hStr := q.Get("height"); hStr != "" {
		if h, err := strconv.Atoi(hStr); err == nil {
			opts.Height = h
		}
	}
	return opts
}

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	animations := engine.GlobalRegistry.List()
	if len(animations) == 0 {
		http.Error(w, "No animations", http.StatusNotFound)
		return
	}
	// simple pseudo random based on time
	idx := int(time.Now().UnixNano()) % len(animations)
	anim := animations[idx]
	// redirect to its endpoint
	http.Redirect(w, r, "/"+engineNormalize(anim.Name()), http.StatusTemporaryRedirect)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	animations := engine.GlobalRegistry.List()
	out := make([]animationResponse, 0, len(animations))
	for _, a := range animations {
		out = append(out, animationResponse{
			Name: a.Name(),
			Category: a.Category(),
			Description: a.Description(),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	help := `HS ASCII LIVE - Terminal Usage
Usage: curl https://hs-ascii-1.onrender.com/<animation>

Examples:
  curl https://hs-ascii-1.onrender.com/pacman
  curl https://hs-ascii-1.onrender.com/matrix-rain?speed=fast&color=green
  curl https://hs-ascii-1.onrender.com/random

Endpoints:
  /<animation>        Stream animation
  /random             Random animation
  /list               List animations
  /help               This help
  /health             Health check
  /api/animations     JSON catalogue
  /api/stats          Usage stats

Parameters:
  speed=slow|normal|fast
  fps=1-60
  color=red|green|yellow|blue|magenta|cyan|white
  width, height
`
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(help))
}

func engineNormalize(s string) string {
	out := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			out += string(r+32)
		} else {
			out += string(r)
		}
	}
	return out
}
