package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"evergon/internal/config"
	"evergon/internal/services"
	"evergon/internal/util"
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

// ---------- PROJECT RUNTIME CLEANUP ----------

func (e *Engine) cleanupProjectRuntimes() {
	runtimeDir := filepath.Join(e.BasePath, "runtime")

	entries, err := os.ReadDir(runtimeDir)
	if err != nil {
		// runtime belum ada, gak apa-apa
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		project := entry.Name()
		projectRuntime := filepath.Join(runtimeDir, project)
		runDir := filepath.Join(projectRuntime, "run")

		phpPid := filepath.Join(runDir, "php-fpm.pid")
		nginxPid := filepath.Join(runDir, "nginx.pid")

		// kill php-fpm by PID
		killPidFile(phpPid)

		// kill nginx by PID
		killPidFile(nginxPid)

		// extra safety: kill by port (dari config project, kalau ada)
		cfg, err := LoadProjectConfig(e.BasePath, project)
		if err == nil {
			// port nginx project
			killPort(cfg.Port)
			// port php-fpm project (port+100)
			killPort(cfg.Port + 100)
		}
	}
}

// ---------- ENGINE START / STOP ----------
// Mulai seluruh service
func (e *Engine) StartAll() error {
	fmt.Println("=== EVERGON START ===")

	pidFile := filepath.Join(e.BasePath, "runtime", "evergon.pid")

	if pid, err := ReadPID(pidFile); err == nil {
		if util.IsAlive(pid) {
			fmt.Println("Evergon already running.")
			return nil
		}
		fmt.Println("Found stale Evergon PID, cleaning up ...")
		KillPID(pidFile)
	}

	// tulis PID engine utama
	if err := os.MkdirAll(filepath.Dir(pidFile), 0o755); err != nil {
		return fmt.Errorf("failed to create runtime dir: %w", err)
	}
	mainPID := os.Getpid()
	if err := WritePID(pidFile, mainPID); err != nil {
		return fmt.Errorf("failed to write evergon pid: %w", err)
	}

	// ============================================
	// AUTO HOSTS GENERATOR (SCAN /WWW → ADD DOMAINS)
	// ============================================
	wwwDir := filepath.Join(e.BasePath, "www")
	entries, _ := os.ReadDir(wwwDir)

	var domains []string
	for _, entry := range entries {
		if entry.IsDir() {
			domains = append(domains, entry.Name()+".test")
		}
	}

	if len(domains) > 0 {
		fmt.Println("[Hosts] Syncing domains...")
		if err := util.EnsureHosts(domains); err != nil {
			fmt.Println("[Hosts] Failed to update /etc/hosts:", err)
		}
	}
	// ============================================

	// start global services
	for _, s := range e.Services {
		fmt.Println("Starting:", s.Name())
		if err := s.Start(); err != nil {
			return err
		}
	}

	fmt.Println("Using BasePath:", e.BasePath)
	fmt.Println("Scanning WWW:", filepath.Join(e.BasePath, "www"))

	fmt.Println("Evergon running at http://localhost:8080")
	return nil
}

func (e *Engine) StopAll() error {
	fmt.Println("=== EVERGON STOP ===")

	// stop global services
	for _, s := range e.Services {
		fmt.Println("Stopping:", s.Name())
		_ = s.Stop()
	}

	// kill project runtimes
	fmt.Println("[ForceKill] Cleaning all project runtimes ...")
	KillAllProjectRuntimes(e.BasePath)

	// kill Evergon main process (jika dipanggil dari luar)
	mainPIDFile := filepath.Join(e.BasePath, "runtime", "evergon.pid")
	KillPID(mainPIDFile)

	fmt.Println("Evergon stopped cleanly.")
	return nil
}

// ---------- STATUS ----------

// Return status tiap service (running / stopped)
func (e *Engine) ServiceStatuses() map[string]services.ServiceStatus {
	statuses := make(map[string]services.ServiceStatus)

	for _, s := range e.Services {
		statuses[s.Name()] = s.Status()
	}

	return statuses
}

// ---------- HELPERS (PID & PORT) ----------

func killPidFile(pidFile string) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return
	}

	pidStr := strings.TrimSpace(string(data))
	if pidStr == "" {
		_ = os.Remove(pidFile)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		_ = os.Remove(pidFile)
		return
	}

	proc, err := os.FindProcess(pid)
	if err == nil {
		_ = proc.Kill()
	}

	_ = os.Remove(pidFile)
}

// kill proses berdasarkan port (fallback paling brutal)
func killPort(port int) {
	out, err := exec.Command("lsof", "-t", "-i", fmt.Sprintf(":%d", port)).Output()
	if err != nil {
		return
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		_ = exec.Command("kill", "-9", line).Run()
	}
}

// ========== PUBLIC: FORCE CLEANUP (dipakai oleh main.go/API) ==========

// Membersihkan semua runtime project secara brutal (PID + PORT)
func (e *Engine) ForceKillAllProjectRuntimes() {
	fmt.Println("[ForceKill] Cleaning all project runtimes ...")

	// 1) cleanup normal
	e.cleanupProjectRuntimes()

	// 2) kill by scanning all project configs
	projectsDir := filepath.Join(e.BasePath, "projects")

	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		project := entry.Name()

		cfg, err := LoadProjectConfig(e.BasePath, project)
		if err != nil {
			continue
		}

		// Kill runtime ports
		killPort(cfg.Port)       // nginx port
		killPort(cfg.Port + 100) // php-fpm port

		// Kill leftover PIDs in runtime folder
		runtimeRun := filepath.Join(e.BasePath, "runtime", project, "run")

		killPidFile(filepath.Join(runtimeRun, "php-fpm.pid"))
		killPidFile(filepath.Join(runtimeRun, "nginx.pid"))
	}
}
