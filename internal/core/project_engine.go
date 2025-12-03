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
	ProjectRoot string // path ke public milik user
	RuntimeRoot string // path ke runtime/{project}
	Services    []services.Service
}

// -----------------------------------------------------------
// INIT
// -----------------------------------------------------------
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

	// Build runtime folders
	for _, dir := range []string{
		runtimeRoot,
		runDir,
		phpDir,
		nginxRoot,
		nginxLogs,
	} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("failed to create runtime folder: %w", err)
		}
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

// -----------------------------------------------------------
// START
// -----------------------------------------------------------
func (e *ProjectEngine) Start() error {
	fmt.Println("=== START PROJECT:", e.Name, "===")

	for _, svc := range e.Services {
		fmt.Println("Starting:", svc.Name())
		if err := svc.Start(); err != nil {
			return fmt.Errorf("error starting %s: %v", svc.Name(), err)
		}
	}

	return nil
}

// -----------------------------------------------------------
// STOP — normal stop
// -----------------------------------------------------------
func (e *ProjectEngine) Stop() error {
	fmt.Println("=== STOP PROJECT:", e.Name, "===")

	for _, svc := range e.Services {
		fmt.Println("Stopping:", svc.Name())
		_ = svc.Stop()
	}

	// FINAL ensure cleanup
	e.cleanupPorts()

	return nil
}

func (e *ProjectEngine) cleanupPorts() {
	phpPort := e.Config.Port + 100
	nginxPort := e.Config.Port

	exec.Command("bash", "-c", fmt.Sprintf("lsof -t -i:%d | xargs -r kill -9", phpPort)).Run()
	exec.Command("bash", "-c", fmt.Sprintf("lsof -t -i:%d | xargs -r kill -9", nginxPort)).Run()
}

// -----------------------------------------------------------
// FORCE STOP — kill ALL related processes brutally
// -----------------------------------------------------------
func (e *ProjectEngine) ForceStopAll() {
	fmt.Println("=== FORCE STOP PROJECT:", e.Name, "===")

	for _, svc := range e.Services {
		_ = svc.Stop()
	}

	e.KillAllProcesses()
}

// -----------------------------------------------------------
// KILL leftover processes: php-fpm workers + nginx
// -----------------------------------------------------------
func (e *ProjectEngine) KillAllProcesses() {
	project := e.Name

	// Kill php-fpm workers that contain project name
	out, _ := exec.Command("bash", "-c",
		fmt.Sprintf("ps aux | grep php-fpm | grep '%s' | awk '{print $2}'", project),
	).Output()

	for _, pid := range strings.Split(string(out), "\n") {
		pid = strings.TrimSpace(pid)
		if pid != "" {
			exec.Command("kill", "-9", pid).Run()
		}
	}

	// Kill nginx workers that contain project runtime path
	out, _ = exec.Command("bash", "-c",
		fmt.Sprintf("ps aux | grep nginx | grep '%s' | awk '{print $2}'", project),
	).Output()

	for _, pid := range strings.Split(string(out), "\n") {
		pid = strings.TrimSpace(pid)
		if pid != "" {
			exec.Command("kill", "-9", pid).Run()
		}
	}
}

// -----------------------------------------------------------
// CLEANUP runtime dirs if needed
// -----------------------------------------------------------
func (e *ProjectEngine) CleanRuntime() {
	os.RemoveAll(e.RuntimeRoot)
}

// -----------------------------------------------------------
// STATUS
// -----------------------------------------------------------
func (e *ProjectEngine) Status() map[string]services.ServiceStatus {
	resp := map[string]services.ServiceStatus{}
	for _, svc := range e.Services {
		resp[svc.Name()] = svc.Status()
	}
	return resp
}
