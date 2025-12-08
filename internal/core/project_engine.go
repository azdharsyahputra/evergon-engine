package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"evergon/internal/services"
)

type ProjectEngine struct {
	BasePath    string
	Name        string
	Config      *ProjectConfig
	ProjectRoot string
	RuntimeRoot string
	Services    []services.Service
}

// ------------------------------------------------------------
// INIT
// ------------------------------------------------------------
func NewProjectEngine(base, name string) (*ProjectEngine, error) {

	cfg, err := LoadProjectConfig(base, name)
	if err != nil {
		return nil, err
	}

	projectRoot := filepath.Join(base, "projects", name)
	publicDir := filepath.Join(projectRoot, cfg.Root)

	runtimeRoot := filepath.Join(base, "runtime", name)
	runDir := filepath.Join(runtimeRoot, "run")
	phpDir := filepath.Join(runtimeRoot, "php")
	nginxRoot := filepath.Join(runtimeRoot, "nginx")
	nginxLogs := filepath.Join(runtimeRoot, "nginx/logs")

	for _, dir := range []string{runtimeRoot, runDir, phpDir, nginxRoot, nginxLogs} {
		_ = os.MkdirAll(dir, 0o755)
	}

	e := &ProjectEngine{
		BasePath:    base,
		Name:        name,
		Config:      cfg,
		ProjectRoot: publicDir,
		RuntimeRoot: runtimeRoot,
	}

	e.Services = []services.Service{
		services.NewProjectPHPService(base, name, cfg.PHPVersion, cfg.Port+100),
		services.NewProjectNginxService(base, name, cfg.Port),
	}

	return e, nil
}

// ------------------------------------------------------------
// START
// ------------------------------------------------------------
func (e *ProjectEngine) Start() error {
	fmt.Println("=== START PROJECT:", e.Name, "===")

	// Start PHP first → then Nginx
	if err := e.Services[0].Start(); err != nil {
		return fmt.Errorf("error starting PHP: %w", err)
	}
	if err := e.Services[1].Start(); err != nil {
		return fmt.Errorf("error starting Nginx: %w", err)
	}

	return nil
}

// ------------------------------------------------------------
// STOP
// ------------------------------------------------------------
func (e *ProjectEngine) Stop() error {
	fmt.Println("=== STOP PROJECT:", e.Name, "===")

	// Stop PHP pool first
	_ = e.Services[0].Stop()

	// Stop Nginx-project second
	_ = e.Services[1].Stop()

	// Cleanup leftover processes
	e.killLeftovers()

	return nil
}

// ------------------------------------------------------------
// KILL leftovers
// ------------------------------------------------------------
func (e *ProjectEngine) killLeftovers() {

	// Kill php-fpm workers for this pool
	out, _ := exec.Command("bash", "-c",
		fmt.Sprintf("ps aux | grep 'php-fpm: pool evergon_%s' | awk '{print $2}'", e.Name),
	).Output()

	for _, pid := range strings.Split(string(out), "\n") {
		p := strings.TrimSpace(pid)
		if p != "" {
			exec.Command("kill", "-9", p).Run()
		}
	}

	// Kill nginx workers tied to this project
	out, _ = exec.Command("bash", "-c",
		fmt.Sprintf("ps aux | grep nginx | grep '%s/runtime/%s' | awk '{print $2}'",
			e.BasePath, e.Name),
	).Output()

	for _, pid := range strings.Split(string(out), "\n") {
		p := strings.TrimSpace(pid)
		if p != "" {
			exec.Command("kill", "-9", p).Run()
		}
	}
}

// ------------------------------------------------------------
// STATUS
// ------------------------------------------------------------
func (e *ProjectEngine) Status() map[string]services.ServiceStatus {
	resp := map[string]services.ServiceStatus{}
	for _, svc := range e.Services {
		resp[svc.Name()] = svc.Status()
	}
	return resp
}
