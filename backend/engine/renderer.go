package engine

import (
	"fmt"
	"time"
)

func RenderFrame(frame Frame, opts AnimationOptions) string {
	// Home cursor to (1,1) so we overwrite in-place without clearing flicker
	home := "\x1b[H"
	content := frame.Content

	// Apply color if requested
	if opts.Color != "" {
		colorCode := colorToCode(opts.Color)
		if colorCode != "" {
			content = fmt.Sprintf("\x1b[%sm%s\x1b[0m", colorCode, content)
		}
	}

	return home + content
}

func colorToCode(name string) string {
	switch name {
	case "red":
		return "31"
	case "green":
		return "32"
	case "yellow":
		return "33"
	case "blue":
		return "34"
	case "magenta":
		return "35"
	case "cyan":
		return "36"
	case "white":
		return "37"
	default:
		return ""
	}
}

func SleepForFrame(frame Frame, opts AnimationOptions) time.Duration {
	fps := opts.FPS
	if fps <= 0 {
		fps = 10
	}
	multiplier := 1.0
	switch opts.Speed {
	case "slow":
		multiplier = 2.0
	case "fast":
		multiplier = 0.5
	}
	base := time.Second / time.Duration(fps)
	d := time.Duration(float64(base) * multiplier)
	if frame.Delay > 0 {
		d = time.Duration(frame.Delay) * time.Millisecond
	}
	return d
}
