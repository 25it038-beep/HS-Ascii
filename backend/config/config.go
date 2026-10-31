package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port                string
	MaxConcurrentClients int
	RateLimitPerMinute  int
	DefaultFPS          int
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	maxClients, _ := strconv.Atoi(getEnv("MAX_CONCURRENT_CLIENTS", "100"))
	rateLimit, _ := strconv.Atoi(getEnv("RATE_LIMIT_PER_MINUTE", "60"))
	fps, _ := strconv.Atoi(getEnv("DEFAULT_FPS", "10"))

	return &Config{
		Port:                 port,
		MaxConcurrentClients: maxClients,
		RateLimitPerMinute:  rateLimit,
		DefaultFPS:          fps,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
