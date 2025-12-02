package api

import (
	"encoding/json"
	"net/http"

	"evergon/internal/core"
)

func RegisterRoutes(mux *http.ServeMux, engine *core.Engine) {
	mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, engine.ServiceStatuses())
	})

	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		_ = engine.StartAll()
		writeJSON(w, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		_ = engine.StopAll()
		writeJSON(w, map[string]string{"status": "ok"})
	})
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
