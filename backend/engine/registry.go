package engine

import (
	"sync"
)

type Registry struct {
	mu         sync.RWMutex
	animations map[string]Animation
}

func NewRegistry() *Registry {
	return &Registry{
		animations: make(map[string]Animation),
	}
}

func (r *Registry) Register(anim Animation) {
	r.mu.Lock()
	defer r.mu.Unlock()
	key := normalize(anim.Name())
	r.animations[key] = anim
}

func (r *Registry) Get(name string) (Animation, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	anim, ok := r.animations[normalize(name)]
	return anim, ok
}

func (r *Registry) List() []Animation {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Animation, 0, len(r.animations))
	for _, a := range r.animations {
		out = append(out, a)
	}
	return out
}

func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.animations))
	for k := range r.animations {
		names = append(names, k)
	}
	return names
}

func normalize(s string) string {
	// simple lower case
	// no external deps
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			out = append(out, r+32)
		} else {
			out = append(out, r)
		}
	}
	return string(out)
}
