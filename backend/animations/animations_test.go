package animations

import (
	"hs-ascii-live/engine"
	"strings"
	"testing"
)

func TestRegisterAll(t *testing.T) {
	r := engine.NewRegistry()
	RegisterAll(r)

	anims := r.List()
	if len(anims) < 180 {
		t.Fatalf("expected at least 180 animations, got %d", len(anims))
	}

	categories := make(map[string]int)
	for _, a := range anims {
		if a.Name() == "" {
			t.Errorf("animation has empty name")
		}
		if a.Category() == "" {
			t.Errorf("animation %s has empty category", a.Name())
		}
		if a.Description() == "" {
			t.Errorf("animation %s has empty description", a.Name())
		}
		categories[a.Category()]++

		// Test frame rendering for multiple frames
		for _, frameNum := range []int{0, 1, 5, 12, 45, 99} {
			frame := a.NextFrame(engine.AnimationOptions{
				Width:  40,
				Height: 15,
			}, frameNum)

			if frame.Content == "" {
				t.Fatalf("animation %s produced empty frame %d", a.Name(), frameNum)
			}
			if !strings.Contains(frame.Content, "HS ASCII LIVE") {
				t.Errorf("animation %s frame %d missing header", a.Name(), frameNum)
			}
		}

		// Test small and extreme dimensions
		smallFrame := a.NextFrame(engine.AnimationOptions{
			Width:  5,
			Height: 3,
		}, 0)
		if smallFrame.Content == "" {
			t.Errorf("animation %s produced empty content for small frame", a.Name())
		}

		largeFrame := a.NextFrame(engine.AnimationOptions{
			Width:  80,
			Height: 30,
		}, 3)
		if largeFrame.Content == "" {
			t.Errorf("animation %s produced empty content for large frame", a.Name())
		}
	}

	expectedCategories := []string{
		"characters", "space", "nature", "technology", "abstract",
		"symbols", "hsai", "gaming", "vehicles", "creatures",
	}

	for _, cat := range expectedCategories {
		count, ok := categories[cat]
		if !ok || count == 0 {
			t.Errorf("category %s has no registered animations", cat)
		}
	}
}

func TestPatterns(t *testing.T) {
	patterns := []string{
		"pacman", "cat", "dancer", "ghost", "space-invaders", "pong", "tetris", "snake",
		"cube", "donut", "dna", "earth", "spiral", "matrix", "equalizer", "rocket",
		"neural", "radar", "blackhole", "vehicle", "hsai", "wings", "symbol",
		"heart", "fire", "waves", "rain", "particles", "stars", "clock", "rotate", "bounce",
	}

	for _, p := range patterns {
		content := generatePattern(p, 40, 15, 0, 123, p)
		if content == "" {
			t.Errorf("pattern %s generated empty content", p)
		}
		if !strings.Contains(content, strings.ToUpper(p)) {
			t.Errorf("pattern %s title missing from header", p)
		}
	}
}
