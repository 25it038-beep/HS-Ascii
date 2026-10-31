package engine

import "testing"

type dummyAnim struct{}
func (d dummyAnim) Name() string { return "test" }
func (d dummyAnim) Category() string { return "test" }
func (d dummyAnim) Description() string { return "test" }
func (d dummyAnim) NextFrame(opts AnimationOptions, frameNum int) Frame { return Frame{Content:"x"} }

func TestRegistry(t *testing.T) {
	r := NewRegistry()
	r.Register(dummyAnim{})
	a, ok := r.Get("test")
	if !ok {
		t.Fatal("expected anim")
	}
	if a.Name() != "test" {
		t.Fatal("name mismatch")
	}
}
