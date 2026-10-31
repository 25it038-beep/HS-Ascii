package engine

type Frame struct {
	Content string
	Delay   int // milliseconds
}

type AnimationOptions struct {
	Speed string // slow, normal, fast
	FPS   int
	Color string
	Width int
	Height int
}

type Animation interface {
	Name() string
	Category() string
	Description() string
	NextFrame(opts AnimationOptions, frameNum int) Frame
}
