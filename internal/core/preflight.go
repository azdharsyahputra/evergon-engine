package core

import (
	"os"
	"os/exec"
	"path/filepath"
)

type CheckResult struct {
	Name   string
	OK     bool
	Reason string
	Fix    string
}

func (e *Engine) PreflightChecks() []CheckResult {
	var results []CheckResult

	// 1. nginx binary
	nginxBin := filepath.Join(e.BasePath, "nginx/sbin/nginx")
	if _, err := os.Stat(nginxBin); err != nil {
		results = append(results, CheckResult{
			Name:   "Nginx binary",
			OK:     false,
			Reason: "nginx binary not found",
		})
	} else {
		results = append(results, CheckResult{Name: "Nginx binary", OK: true})
	}

	// 2. nginx capability (port 80)
	if !hasCap(nginxBin) {
		results = append(results, CheckResult{
			Name:   "Nginx trust",
			OK:     false,
			Reason: "port 80 requires privileged permission",
			Fix:    "evergon setup",
		})
	} else {
		results = append(results, CheckResult{Name: "Nginx trust", OK: true})
	}

	// 3. PHP-FPM
	if !processRunning("php-fpm") {
		results = append(results, CheckResult{
			Name:   "PHP-FPM",
			OK:     false,
			Reason: "php-fpm not running",
		})
	} else {
		results = append(results, CheckResult{Name: "PHP-FPM", OK: true})
	}

	return results
}

func processRunning(name string) bool {
	err := exec.Command("pgrep", name).Run()
	return err == nil
}
