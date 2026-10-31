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
		Speed:  "normal",
		FPS:    10,
		Width:  80,
		Height: 24,
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

	// Fullscreen flags for large/maximized terminals
	if isFull := q.Get("fullscreen"); isFull == "true" || isFull == "1" {
		opts.Width = 120
		opts.Height = 35
	} else if isFull := q.Get("full"); isFull == "true" || isFull == "1" {
		opts.Width = 120
		opts.Height = 35
	} else if size := q.Get("size"); size == "full" || size == "max" {
		opts.Width = 120
		opts.Height = 35
	} else if size == "small" {
		opts.Width = 40
		opts.Height = 15
	}

	// Width aliases: width, w, cols, columns
	wStr := q.Get("width")
	if wStr == "" {
		wStr = q.Get("w")
	}
	if wStr == "" {
		wStr = q.Get("cols")
	}
	if wStr == "" {
		wStr = q.Get("columns")
	}
	if wStr != "" {
		if w, err := strconv.Atoi(wStr); err == nil && w > 0 {
			opts.Width = w
		}
	}

	// Height aliases: height, h, lines, rows
	hStr := q.Get("height")
	if hStr == "" {
		hStr = q.Get("h")
	}
	if hStr == "" {
		hStr = q.Get("lines")
	}
	if hStr == "" {
		hStr = q.Get("rows")
	}
	if hStr != "" {
		if h, err := strconv.Atoi(hStr); err == nil && h > 0 {
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
  curl https://hs-ascii-1.onrender.com/matrix-rain?color=green
  curl https://hs-ascii-1.onrender.com/space-invaders?full=true
  curl https://hs-ascii-1.onrender.com/rotating-cube?w=120&h=35
  curl https://hs-ascii-1.onrender.com/random

Endpoints:
  /<animation>        Stream continuous animation
  /random             Random animation
  /list               List animations
  /help               This help
  /health             Health check
  /api/animations     JSON catalogue
  /api/stats          Usage stats

Screen & Sizing Parameters:
  full=true           Maximized terminal mode (120x35)
  size=full|small     full (120x35) or small (40x15) [default: 80x24 standard full CMD]
  w, width, cols      Custom terminal width in columns
  h, height, rows     Custom terminal height in rows

Animation Parameters:
  speed=slow|normal|fast
  fps=1-60
  color=red|green|yellow|blue|magenta|cyan|white
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
