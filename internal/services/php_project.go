package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type ProjectPHPService struct {
	BasePath string
	Project  string
	Version  string
}

// ------------------------------------------------------------
// HELPERS
// ------------------------------------------------------------

func (s *ProjectPHPService) Name() string {
	return "php-pool:" + s.Project
}

func (s *ProjectPHPService) poolName() string {
	return "evergon_" + strings.ReplaceAll(s.Project, "-", "_")
}

func (s *ProjectPHPService) phpBase() string {
	return filepath.Join(s.BasePath, "php", s.Version)
}

func (s *ProjectPHPService) poolConfPath() string {
	return filepath.Join(s.phpBase(), "etc", "php-fpm.d", s.poolName()+".conf")
}

func (s *ProjectPHPService) sockPath() string {
	return filepath.Join(s.BasePath, "runtime", s.Project, "php", "php-fpm.sock")
}

func (s *ProjectPHPService) logPath() string {
	return filepath.Join(s.BasePath, "runtime", s.Project, "logs", "php-fpm.log")
}

func (s *ProjectPHPService) globalPidFile() string {
	// pid file FPM global (yang valid di portable PHP-FPM)
	return filepath.Join(s.phpBase(), "logs", "php-fpm.pid")
}

// ------------------------------------------------------------
// START — create pool & reload FPM
// ------------------------------------------------------------

func (s *ProjectPHPService) Start() error {

	// Ensure dirs
	runtimePHP := filepath.Join(s.BasePath, "runtime", s.Project, "php")
	runtimeLog := filepath.Join(s.BasePath, "runtime", s.Project, "logs")
	_ = os.MkdirAll(runtimePHP, 0o755)
	_ = os.MkdirAll(runtimeLog, 0o755)

	// Build pool config
	poolConf := fmt.Sprintf(`
[%s]
listen = %s
listen.owner = %s
listen.group = %s
listen.mode = 0660

pm = dynamic
pm.max_children = 5
pm.start_servers = 2
pm.min_spare_servers = 1
pm.max_spare_servers = 3

php_admin_value[error_log] = %s
php_admin_flag[log_errors] = on
`, s.poolName(), s.sockPath(), os.Getenv("USER"), os.Getenv("USER"), s.logPath())

	if err := os.WriteFile(s.poolConfPath(), []byte(poolConf), 0o644); err != nil {
		return fmt.Errorf("failed writing pool file: %w", err)
	}

	// Reload FPM to activate new pool
	return s.reloadFPM()
}

// ------------------------------------------------------------
// STOP — remove pool & reload FPM (twice to flush workers)
// ------------------------------------------------------------

func (s *ProjectPHPService) Stop() error {
	// Remove pool config
	_ = os.Remove(s.poolConfPath())

	// Remove socket
	_ = os.Remove(s.sockPath())

	// Reload FPM twice → untuk flush worker pool
	_ = s.reloadFPM()
	_ = s.reloadFPM()

	return nil
}

// ------------------------------------------------------------
// STATUS
// ------------------------------------------------------------

func (s *ProjectPHPService) Status() ServiceStatus {
	if _, err := os.Stat(s.sockPath()); err == nil {
		return ServiceStatus{Running: true}
	}
	return ServiceStatus{Running: false}
}

// ------------------------------------------------------------
// INTERNAL — reload php-fpm master
// ------------------------------------------------------------

func (s *ProjectPHPService) reloadFPM() error {
	data, err := os.ReadFile(s.globalPidFile())
	if err != nil {
		return fmt.Errorf("php-fpm not running for version %s", s.Version)
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return err
	}

	// USR2 = soft reload (safe)
	cmd := exec.Command("kill", "-USR2", strconv.Itoa(pid))
	return cmd.Run()
}
