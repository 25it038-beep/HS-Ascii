package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"hs-ascii-live/animations"
	"hs-ascii-live/engine"

	"github.com/gorilla/mux"
)

func init() {
	animations.RegisterAll(engine.GlobalRegistry)
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	HealthHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHelpHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/help", nil)
	w := httptest.NewRecorder()
	HelpHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	body := w.Body.String()
	if !strings.Contains(body, "HS ASCII LIVE") {
		t.Fatalf("expected help text to contain HS ASCII LIVE")
	}
}

func TestListHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/list", nil)
	w := httptest.NewRecorder()
	ListHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var list []animationResponse
	if err := json.Unmarshal(w.Body.Bytes(), &list); err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}
	if len(list) < 180 {
		t.Fatalf("expected >= 180 animations in /list, got %d", len(list))
	}
}

func TestAnimationsAPI(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/animations", nil)
	w := httptest.NewRecorder()
	AnimationsAPI(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var list []animMeta
	if err := json.Unmarshal(w.Body.Bytes(), &list); err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}
	if len(list) < 180 {
		t.Fatalf("expected >= 180 animations in /api/animations, got %d", len(list))
	}
}

func TestStatsAPI(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/stats", nil)
	w := httptest.NewRecorder()
	StatsAPI(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var stats map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &stats); err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}
	total, ok := stats["total_animations"].(float64)
	if !ok || int(total) < 180 {
		t.Fatalf("expected total_animations >= 180, got %v", stats["total_animations"])
	}
}

func TestRandomHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/random", nil)
	w := httptest.NewRecorder()
	RandomHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusTemporaryRedirect {
		t.Fatalf("expected redirect 307, got %d", resp.StatusCode)
	}
	loc := resp.Header.Get("Location")
	if loc == "" || loc == "/" {
		t.Fatalf("invalid redirect location: %s", loc)
	}
}

func TestAnimationHandlerStreaming(t *testing.T) {
	testCases := []string{
		"pacman",
		"matrix-rain",
		"space-invaders",
		"rotating-cube",
		"cybertruck",
		"dragon-flight",
		"hs-ai-solutions-logo",
	}

	for _, name := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		req := httptest.NewRequest("GET", "/"+name+"?fps=30", nil).WithContext(ctx)
		req = mux.SetURLVars(req, map[string]string{"name": name})

		w := httptest.NewRecorder()
		AnimationHandler(w, req)

		body := w.Body.String()
		if !strings.Contains(body, "HS ASCII LIVE") {
			t.Errorf("animation %s stream missing header", name)
		}
	}
}

func TestAnimationHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/non-existent-animation-xyz", nil)
	req = mux.SetURLVars(req, map[string]string{"name": "non-existent-animation-xyz"})
	w := httptest.NewRecorder()
	AnimationHandler(w, req)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404 for unknown animation, got %d", w.Result().StatusCode)
	}
}
