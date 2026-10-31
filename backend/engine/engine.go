package engine

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var GlobalRegistry = NewRegistry()

func StreamAnimation(ctx context.Context, w io.Writer, anim Animation, opts AnimationOptions) error {
	// Initial clear screen and hide cursor (\x1b[?25l)
	_, _ = fmt.Fprint(w, "\x1b[?25l\x1b[2J\x1b[H")
	if rw, ok := w.(http.ResponseWriter); ok {
		rc := http.NewResponseController(rw)
		_ = rc.Flush()
	}

	defer func() {
		// Restore cursor (\x1b[?25h) on disconnect
		_, _ = fmt.Fprint(w, "\x1b[?25h\r\n")
		if rw, ok := w.(http.ResponseWriter); ok {
			rc := http.NewResponseController(rw)
			_ = rc.Flush()
		}
	}()

	frameNum := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		frame := anim.NextFrame(opts, frameNum)
		out := RenderFrame(frame, opts)
		_, err := fmt.Fprint(w, out)
		if err != nil {
			return err
		}
		// flush immediately so streaming works through reverse proxies & Cloudflare
		if rw, ok := w.(http.ResponseWriter); ok {
			rc := http.NewResponseController(rw)
			_ = rc.Flush()
		} else if f, ok := w.(http.Flusher); ok {
			f.Flush()
		} else if f, ok := w.(interface{ Flush() }); ok {
			f.Flush()
		}

		delay := SleepForFrame(frame, opts)
		frameNum++
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
