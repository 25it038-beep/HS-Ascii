package handlers

import (
	"encoding/json"
	"hs-ascii-live/engine"
	"net/http"
)

type animMeta struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func AnimationsAPI(w http.ResponseWriter, r *http.Request) {
	animations := engine.GlobalRegistry.List()
	out := make([]animMeta, 0, len(animations))
	for _, a := range animations {
		out = append(out, animMeta{
			Name: a.Name(),
			Category: a.Category(),
			Description: a.Description(),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func StatsAPI(w http.ResponseWriter, r *http.Request) {
	stats := map[string]interface{}{
		"total_animations": len(engine.GlobalRegistry.List()),
		"status": "ok",
		"service": "hs-ascii-live",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
