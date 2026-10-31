package engine

import (
	"context"
	"fmt"
	"io"
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
		// flush if possible
		if f, ok := w.(interface{ Flush() error }); ok {
			_ = f.Flush()
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
