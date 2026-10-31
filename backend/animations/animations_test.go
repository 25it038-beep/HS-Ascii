package animations

import (
	"fmt"
	"hs-ascii-live/engine"
	"strings"
	"testing"
)

func TestRegisterAll(t *testing.T) {
	r := engine.NewRegistry()
	RegisterAll(r)

	anims := r.List()
	t.Logf("TOTAL REGISTERED ANIMATIONS: %d", len(anims))
	if len(anims) < 300 {
		t.Fatalf("expected at least 300 animations, got %d", len(anims))
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
				Width:  80,
				Height: 24,
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
			Width:  120,
			Height: 35,
		}, 3)
		if largeFrame.Content == "" {
			t.Errorf("animation %s produced empty content for large frame", a.Name())
		}
	}

	expectedCategories := []string{
		"characters", "space", "nature", "technology", "abstract",
		"symbols", "hsai", "gaming", "vehicles", "creatures",
		"music", "sports", "cyberpunk", "food", "weather",
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
		"piano", "vinyl", "city", "coffee", "sports", "tux", "pokeball", "tornado", "disco",
	}

	for _, p := range patterns {
		content := generatePattern(p, 80, 24, 0, 123, p)
		if content == "" {
			t.Errorf("pattern %s generated empty content", p)
		}
		if !strings.Contains(content, strings.ToUpper(p)) {
			t.Errorf("pattern %s title missing from header", p)
		}
	}
}

func TestFindDuplicates(t *testing.T) {
	r := engine.NewRegistry()
	RegisterAll(r)

	anims := r.List()
	t.Logf("Total animations: %d", len(anims))

	// Check name uniqueness
	nameSeen := make(map[string]bool)
	var exactDups []string
	for _, a := range anims {
		if nameSeen[a.Name()] {
			exactDups = append(exactDups, a.Name())
		}
		nameSeen[a.Name()] = true
	}
	if len(exactDups) > 0 {
		t.Logf("EXACT DUPLICATE NAMES (%d): %v", len(exactDups), exactDups)
	}

	// Check visual frame content uniqueness (strip header line)
	contentMap := make(map[string][]string)
	opts := engine.AnimationOptions{Width: 80, Height: 24}
	for _, a := range anims {
		f0 := a.NextFrame(opts, 0).Content
		f5 := a.NextFrame(opts, 5).Content
		// Strip first line (which is the title header)
		lines0 := strings.Split(f0, "\r\n")
		lines5 := strings.Split(f5, "\r\n")
		body0 := ""
		body5 := ""
		if len(lines0) > 1 {
			body0 = strings.Join(lines0[1:], "\n")
		}
		if len(lines5) > 1 {
			body5 = strings.Join(lines5[1:], "\n")
		}
		sig := fmt.Sprintf("%s|||%s", body0, body5)
		contentMap[sig] = append(contentMap[sig], a.Name())
	}

	totalVisualDups := 0
	for sig, names := range contentMap {
		_ = sig
		if len(names) > 1 {
			totalVisualDups += len(names) - 1
			t.Logf("VISUAL DUPLICATES (%d sharing same art): %v", len(names), names)
		}
	}
	t.Logf("TOTAL VISUAL DUPLICATES: %d", totalVisualDups)
}
