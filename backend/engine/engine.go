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
