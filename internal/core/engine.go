package core

import (
	"fmt"
	"os"
	"path/filepath"

	"evergon/internal/config"
	"evergon/internal/services"
)

type Engine struct {
	BasePath string
	Config   config.EngineConfig
	Services []services.Service
}

// Constructor utama engine
func NewEngine(base string) *Engine {
	cfgPath := filepath.Join(base, "config", "engine.json")
	cfg := config.Load(cfgPath)

	www := filepath.Join(base, "www")

	e := &Engine{
		BasePath: base,
		Config:   cfg,
	}

	// Build services awal berdasarkan config
	e.Services = []services.Service{
		services.NewPHPService(base, cfg.PHPVersion),
		services.NewNginxService(base, www),
	}

	return e
}

// Menyimpan config (php_version dll)
func (e *Engine) saveConfig() error {
	cfgPath := filepath.Join(e.BasePath, "config", "engine.json")

	if err := os.MkdirAll(filepath.Dir(cfgPath), 0o755); err != nil {
		return err
	}

	return config.Save(cfgPath, e.Config)
}

// Mulai seluruh service
func (e *Engine) StartAll() error {
	fmt.Println("=== EVERGON START ===")

	for _, s := range e.Services {
		fmt.Println("Starting:", s.Name())
		if err := s.Start(); err != nil {
			fmt.Println("Failed to start", s.Name(), ":", err)
			return err
		}
	}

	fmt.Println("Evergon running at http://localhost:8080")
	return nil
}

// Stop seluruh service
func (e *Engine) StopAll() error {
	fmt.Println("=== EVERGON STOP ===")

	for _, s := range e.Services {
		fmt.Println("Stopping:", s.Name())
		if err := s.Stop(); err != nil {
			fmt.Println("Failed to stop", s.Name(), ":", err)
		}
	}

	fmt.Println("Evergon stopped.")
	return nil
}

// Return status tiap service (running / stopped)
func (e *Engine) ServiceStatuses() map[string]services.ServiceStatus {
	statuses := make(map[string]services.ServiceStatus)

	for _, s := range e.Services {
		statuses[s.Name()] = s.Status()
	}

	return statuses
}
