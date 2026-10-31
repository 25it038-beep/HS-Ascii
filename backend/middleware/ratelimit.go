package middleware

import (
	"net/http"
	"sync"
	"time"
)

type visitor struct {
	count int
	last  time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*visitor),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		rl.mu.Lock()
		v, ok := rl.visitors[ip]
		now := time.Now()
		if !ok || now.Sub(v.last) > rl.window {
			rl.visitors[ip] = &visitor{count: 1, last: now}
			rl.mu.Unlock()
			next.ServeHTTP(w, r)
			return
		}
		if v.count >= rl.limit {
			rl.mu.Unlock()
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		v.count++
		rl.mu.Unlock()
		next.ServeHTTP(w, r)
	})
}
