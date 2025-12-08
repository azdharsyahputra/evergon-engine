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

func NewProjectPHPService(base, project, version string, _ int) *ProjectPHPService {
	return &ProjectPHPService{
		BasePath: base,
		Project:  project,
		Version:  version,
	}
}

func (s *ProjectPHPService) Name() string {
	return "php-pool:" + s.Project
}

func (s *ProjectPHPService) poolName() string {
	return "evergon_" + strings.ReplaceAll(s.Project, "-", "_")
}

func (s *ProjectPHPService) phpBase() string {
	return filepath.Join(s.BasePath, "php", s.Version)
}

func (s *ProjectPHPService) poolDir() string {
	return filepath.Join(s.phpBase(), "etc", "php-fpm.d")
}

func (s *ProjectPHPService) sockPath() string {
	return filepath.Join(s.BasePath, "runtime", s.Project, "php", "php-fpm.sock")
}

func (s *ProjectPHPService) logPath() string {
	return filepath.Join(s.BasePath, "runtime", s.Project, "logs", "php-fpm.log")
}

func (s *ProjectPHPService) pidFile() string {
	return filepath.Join(s.phpBase(), "var", "run", "php-fpm.pid")
}

func (s *ProjectPHPService) Start() error {
	runtimeRoot := filepath.Join(s.BasePath, "runtime", s.Project)
	phpRuntime := filepath.Join(runtimeRoot, "php")
	logDir := filepath.Join(runtimeRoot, "logs")

	_ = os.MkdirAll(phpRuntime, 0o755)
	_ = os.MkdirAll(logDir, 0o755)
	_ = os.MkdirAll(s.poolDir(), 0o755)

	poolFile := filepath.Join(s.poolDir(), s.poolName()+".conf")

	poolConf := fmt.Sprintf(`
[%s]
listen = %s
listen.mode = 0660

pm = dynamic
pm.max_children = 5
pm.start_servers = 2
pm.min_spare_servers = 1
pm.max_spare_servers = 3

php_admin_value[error_log] = %s
php_admin_flag[log_errors] = on
`, s.poolName(), s.sockPath(), s.logPath())

	if err := os.WriteFile(poolFile, []byte(poolConf), 0o644); err != nil {
		return err
	}

	data, err := os.ReadFile(s.pidFile())
	if err != nil {
		return fmt.Errorf("php-fpm not running for version %s", s.Version)
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return err
	}

	_ = exec.Command("kill", "-USR2", strconv.Itoa(pid)).Run()

	return nil
}

func (s *ProjectPHPService) Stop() error {
	poolFile := filepath.Join(s.poolDir(), s.poolName()+".conf")
	_ = os.Remove(poolFile)

	data, err := os.ReadFile(s.pidFile())
	if err == nil {
		if pid, err2 := strconv.Atoi(strings.TrimSpace(string(data))); err2 == nil {
			_ = exec.Command("kill", "-USR2", strconv.Itoa(pid)).Run()
		}
	}

	_ = os.Remove(s.sockPath())

	return nil
}

func (s *ProjectPHPService) Status() ServiceStatus {
	if _, err := os.Stat(s.sockPath()); err == nil {
		return ServiceStatus{
			Running: true,
			Port:    0,
		}
	}
	return ServiceStatus{Running: false}
}
