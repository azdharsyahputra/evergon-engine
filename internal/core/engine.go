package core

import (
	"fmt"
	"path/filepath"

	"evergon/internal/services"
)

type Engine struct {
	BasePath string
	Services []services.Service
}

func NewEngine(base string) *Engine {
	www := filepath.Join(base, "www")

	return &Engine{
		BasePath: base,
		Services: []services.Service{
			services.NewPHPService(base),
			services.NewNginxService(base, www),
		},
	}
}

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

func (e *Engine) ServiceStatuses() map[string]services.ServiceStatus {
	statuses := make(map[string]services.ServiceStatus)
	for _, s := range e.Services {
		statuses[s.Name()] = s.Status()
	}
	return statuses
}
