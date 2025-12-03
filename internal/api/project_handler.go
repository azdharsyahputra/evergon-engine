package api

import (
	"encoding/json"
	"net/http"

	"evergon/internal/core"
)

type ProjectHandler struct {
	BasePath string
	Registry *core.ProjectRegistry
}

func NewProjectHandler(base string) *ProjectHandler {
	return &ProjectHandler{
		BasePath: base,
		Registry: core.NewProjectRegistry(base),
	}
}

func (h *ProjectHandler) Register(mux *http.ServeMux) {

	// ========================
	// LIST PROJECTS
	// ========================
	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		projects, err := h.Registry.List()
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, projects)
	})

	// ========================
	// CREATE PROJECT
	// ========================
	mux.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Name string `json:"name"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)

		if req.Name == "" {
			writeJSON(w, map[string]string{"error": "name required"})
			return
		}

		err := h.Registry.Create(req.Name)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, map[string]string{
			"project": req.Name,
			"status":  "ok",
		})
	})

	// ========================
	// START PROJECT
	// ========================
	mux.HandleFunc("/project/start", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		engine, err := h.Registry.Load(name)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		err = engine.Start()
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, map[string]string{"status": "running"})
	})

	// ========================
	// STOP PROJECT
	// ========================
	mux.HandleFunc("/project/stop", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		engine, err := h.Registry.Load(name)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		err = engine.Stop()
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, map[string]string{"status": "stopped"})
	})

	// ========================
	// RESTART PROJECT
	// ========================
	mux.HandleFunc("/project/restart", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		engine, err := h.Registry.Load(name)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		_ = engine.Stop()
		err = engine.Start()

		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, map[string]string{"status": "restarted"})
	})

	// ========================
	// PROJECT STATUS
	// ========================
	mux.HandleFunc("/project/status", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		engine, err := h.Registry.Load(name)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, engine.Status())
	})

	// ========================
	// READ PROJECT CONFIG
	// ========================
	mux.HandleFunc("/project/config", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		cfg, err := h.Registry.ReadConfig(name)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, cfg)
	})

	// ========================
	// UPDATE PROJECT CONFIG
	// ========================
	mux.HandleFunc("/project/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Query().Get("name")

		var req core.ProjectConfig
		_ = json.NewDecoder(r.Body).Decode(&req)

		err := h.Registry.Update(name, &req)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, map[string]string{"status": "updated"})
	})
}
