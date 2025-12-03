package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type ProjectPHPService struct {
	BasePath string
	Project  string
	Version  string
	Port     int
}

func NewProjectPHPService(base, project, version string, port int) *ProjectPHPService {
	return &ProjectPHPService{
		BasePath: base,
		Project:  project,
		Version:  version,
		Port:     port,
	}
}

func (s *ProjectPHPService) Name() string {
	return "php-project:" + s.Project
}

func (s *ProjectPHPService) Start() error {

	runtime := filepath.Join(s.BasePath, "runtime", s.Project)
	phpRuntime := filepath.Join(runtime, "php")
	poolDir := filepath.Join(phpRuntime, "php-fpm.d")
	runDir := filepath.Join(runtime, "run")
	logDir := filepath.Join(runtime, "logs")

	phpBase := filepath.Join(s.BasePath, "php", s.Version)

	phpBin := filepath.Join(phpBase, "sbin/php-fpm")
	phpIni := filepath.Join(phpBase, "etc/php.ini")

	globalConf := filepath.Join(phpRuntime, "php-fpm.conf")
	poolConf := filepath.Join(poolDir, "www.conf")

	pidFile := filepath.Join(runDir, "php-fpm.pid")
	logFile := filepath.Join(logDir, "php-fpm.log")

	// ensure dirs
	_ = os.MkdirAll(phpRuntime, 0o755)
	_ = os.MkdirAll(poolDir, 0o755)
	_ = os.MkdirAll(runDir, 0o755)
	_ = os.MkdirAll(logDir, 0o755)

	// GLOBAL CONFIG
	global := fmt.Sprintf(`
[global]
pid = %s
error_log = %s
include=%s/*.conf
`, pidFile, logFile, poolDir)

	_ = os.WriteFile(globalConf, []byte(global), 0o644)

	// POOL CONFIG
	pool := fmt.Sprintf(`
[www]
listen = 127.0.0.1:%d
pm = dynamic
pm.max_children = 5
pm.start_servers = 2
pm.min_spare_servers = 1
pm.max_spare_servers = 3
`, s.Port)

	_ = os.WriteFile(poolConf, []byte(pool), 0o644)

	// RUN PHP-FPM
	cmd := exec.Command(phpBin,
		"-p", phpRuntime,
		"-y", globalConf,
		"-c", phpIni,
	)

	cmd.Env = append(os.Environ(),
		"LD_LIBRARY_PATH="+filepath.Join(phpBase, "libs"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (s *ProjectPHPService) Stop() error {
	runtimeRoot := filepath.Join(s.BasePath, "runtime", s.Project)
	pidFile := filepath.Join(runtimeRoot, "run/php-fpm.pid")

	// 1. Kill master
	data, err := os.ReadFile(pidFile)
	if err == nil {
		if pid, err := strconv.Atoi(string(data)); err == nil {
			proc, _ := os.FindProcess(pid)
			_ = proc.Kill()
		}
		_ = os.Remove(pidFile)
	}

	// 2. Kill ALL workers listening on s.Port
	killCmd := fmt.Sprintf(
		"lsof -t -i:%d | xargs -r kill -9",
		s.Port,
	)
	exec.Command("bash", "-c", killCmd).Run()

	return nil
}

func (s *ProjectPHPService) Status() ServiceStatus {
	pidFile := filepath.Join(s.BasePath, "runtime", s.Project, "run/php-fpm.pid")

	data, err := os.ReadFile(pidFile)
	if err != nil {
		return ServiceStatus{Running: false, Port: s.Port}
	}

	pid, _ := strconv.Atoi(string(data))

	return ServiceStatus{
		Running: true,
		PID:     pid,
		Port:    s.Port,
	}
}
