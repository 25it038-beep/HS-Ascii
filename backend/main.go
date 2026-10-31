package main

import (
	"log"
	"net/http"
	"time"

	"hs-ascii-live/animations"
	"hs-ascii-live/config"
	"hs-ascii-live/engine"
	"hs-ascii-live/handlers"
	"hs-ascii-live/middleware"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	// Register animations
	animations.RegisterAll(engine.GlobalRegistry)
	log.Printf("Registered %d animations", len(engine.GlobalRegistry.List()))

	r := mux.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)

	rl := middleware.NewRateLimiter(cfg.RateLimitPerMinute, time.Minute)
	r.Use(rl.Middleware)

	// API routes
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/help", handlers.HelpHandler).Methods("GET")
	r.HandleFunc("/list", handlers.ListHandler).Methods("GET")
	r.HandleFunc("/random", handlers.RandomHandler).Methods("GET")
	r.HandleFunc("/api/animations", handlers.AnimationsAPI).Methods("GET")
	r.HandleFunc("/api/stats", handlers.StatsAPI).Methods("GET")

	// Animation routes
	r.HandleFunc("/{name}", handlers.AnimationHandler).Methods("GET")

	log.Printf("Starting server on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
