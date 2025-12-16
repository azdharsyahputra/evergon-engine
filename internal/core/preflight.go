package core

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

// ------------------------------------
// RESULT MODEL
// ------------------------------------

type CheckResult struct {
	Name   string
	OK     bool
	Reason string
	Fix    string
}

// ------------------------------------
// PREFLIGHT ENTRYPOINT
// ------------------------------------

func (e *Engine) PreflightChecks() []CheckResult {
	var results []CheckResult

	nginxBin := filepath.Join(e.BasePath, "nginx", "sbin", "nginx")

	// 1. nginx binary exists
	if _, err := os.Stat(nginxBin); err != nil {
		results = append(results, CheckResult{
			Name:   "Nginx binary",
			OK:     false,
			Reason: "nginx binary not found",
			Fix:    "reinstall or rebuild Evergon",
		})
	} else {
		results = append(results, CheckResult{
			Name: "Nginx binary",
			OK:   true,
		})
	}

	// 2. nginx trust (cap_net_bind_service)
	if !hasNetBindCap(nginxBin) {
		results = append(results, CheckResult{
			Name:   "Nginx trust",
			OK:     false,
			Reason: "port 80 requires privileged permission",
			Fix:    "run: evergon setup",
		})
	} else {
		results = append(results, CheckResult{
			Name: "Nginx trust",
			OK:   true,
		})
	}

	// 3. port 80 availability
	if !isPortAvailable(80) {
		results = append(results, CheckResult{
			Name:   "Port 80",
			OK:     false,
			Reason: "port 80 is already in use",
			Fix:    "stop existing service or run: evergon stop",
		})
	} else {
		results = append(results, CheckResult{
			Name: "Port 80",
			OK:   true,
		})
	}

	// 4. PHP-FPM running
	if !processRunning("php-fpm") {
		results = append(results, CheckResult{
			Name:   "PHP-FPM",
			OK:     false,
			Reason: "php-fpm is not running",
			Fix:    "start PHP-FPM or check Evergon logs",
		})
	} else {
		results = append(results, CheckResult{
			Name: "PHP-FPM",
			OK:   true,
		})
	}

	return results
}

// ------------------------------------
// HELPERS
// ------------------------------------

// check if nginx has cap_net_bind_service
func hasNetBindCap(bin string) bool {
	out, err := exec.Command("getcap", bin).Output()
	if err != nil {
		return false
	}
	return len(out) > 0
}

// native Go port availability check (no sudo)
func isPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	_ = ln.Close()
	return true
}

// check if a process name exists
func processRunning(name string) bool {
	err := exec.Command("pgrep", name).Run()
	return err == nil
}
