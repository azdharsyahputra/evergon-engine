package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  evergon start")
		fmt.Println("  evergon stop")
		return
	}

	base, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	switch os.Args[1] {
	case "start":
		startEvergon(base)
	case "stop":
		stopEvergon(base)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

////////////////////////////////////////////////////////////////////////////////
// START EVERGON
////////////////////////////////////////////////////////////////////////////////

func startEvergon(base string) {
	fmt.Println("=== EVERGON START ===")

	nginxBase := filepath.Join(base, "nginx")
	phpBase := filepath.Join(base, "php")
	wwwBase := filepath.Join(base, "www")

	// 0. Bersihkan port yang mau dipake Evergon
	killPort(8080)
	killPort(9099)

	// 1. Hapus PID stale
	cleanupStalePID(filepath.Join(nginxBase, "logs/nginx.pid"))
	cleanupStalePID(filepath.Join(phpBase, "logs/php-fpm.pid"))

	// 2. Siapkan folder temp/logs
	prepareNginxDirs(nginxBase)
	preparePHPDirs(phpBase)

	// 3. Generate nginx.conf
	generateNginxConf(
		filepath.Join(nginxBase, "conf/nginx.conf.tpl"),
		filepath.Join(nginxBase, "conf/nginx.conf"),
		wwwBase,
	)

	// 4. Start PHP-FPM
	if err := startPHPFPM(phpBase); err != nil {
		fmt.Println("Failed to start PHP-FPM:", err)
		return
	}

	// 5. Start Nginx
	if err := startNginx(nginxBase); err != nil {
		fmt.Println("Failed to start Nginx:", err)
		return
	}

	fmt.Println("Evergon running at http://localhost:8080")
}

////////////////////////////////////////////////////////////////////////////////
// STOP EVERGON
////////////////////////////////////////////////////////////////////////////////

func stopEvergon(base string) {
	fmt.Println("=== EVERGON STOP ===")

	nginxPID := filepath.Join(base, "nginx/logs/nginx.pid")
	phpPID := filepath.Join(base, "php/logs/php-fpm.pid")

	stopByPID(nginxPID)
	stopByPID(phpPID)

	killPort(8080)
	killPort(9099)

	fmt.Println("Evergon stopped.")
}

////////////////////////////////////////////////////////////////////////////////
// PREP DIRS
////////////////////////////////////////////////////////////////////////////////

func prepareNginxDirs(nginxBase string) {
	folders := []string{
		"client_body_temp",
		"fastcgi_temp",
		"proxy_temp",
		"scgi_temp",
		"uwsgi_temp",
		"logs",
	}

	for _, f := range folders {
		_ = os.MkdirAll(filepath.Join(nginxBase, f), 0o755)
	}

	touch(filepath.Join(nginxBase, "logs/access.log"))
	touch(filepath.Join(nginxBase, "logs/error.log"))
}

func preparePHPDirs(phpBase string) {
	folders := []string{
		"logs",
		"var/run",
		"var/log",
	}

	for _, f := range folders {
		_ = os.MkdirAll(filepath.Join(phpBase, f), 0o755)
	}

	touch(filepath.Join(phpBase, "logs/php-fpm.log"))
}

func touch(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.WriteFile(path, []byte{}, 0o644)
	}
}

////////////////////////////////////////////////////////////////////////////////
// CONFIG GENERATOR
////////////////////////////////////////////////////////////////////////////////

func generateNginxConf(srcTpl, dstConf, rootPath string) {
	data, err := os.ReadFile(srcTpl)
	if err != nil {
		fmt.Println("Failed to read nginx.conf.tpl:", err)
		return
	}

	content := strings.ReplaceAll(string(data), "{{ROOT_PATH}}", rootPath)

	if err := os.WriteFile(dstConf, []byte(content), 0o644); err != nil {
		fmt.Println("Failed to write nginx.conf:", err)
	}
}

////////////////////////////////////////////////////////////////////////////////
// START PROCESSES
////////////////////////////////////////////////////////////////////////////////

func startPHPFPM(phpBase string) error {
	fpmBin := filepath.Join(phpBase, "sbin/php-fpm")
	conf := filepath.Join(phpBase, "etc/php-fpm.conf")
	ini := filepath.Join(phpBase, "etc/php.ini")

	fmt.Println("Starting PHP-FPM...")

	cmd := exec.Command(fpmBin,
		"-p", phpBase,
		"-y", conf,
		"-c", ini,
		"--daemonize",
	)

	cmd.Env = append(os.Environ(),
		"LD_LIBRARY_PATH="+filepath.Join(phpBase, "libs")+":"+os.Getenv("LD_LIBRARY_PATH"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func startNginx(nginxBase string) error {
	nginxBin := filepath.Join(nginxBase, "sbin/nginx")
	conf := filepath.Join(nginxBase, "conf/nginx.conf")

	fmt.Println("Starting Nginx...")

	cmd := exec.Command(nginxBin,
		"-p", nginxBase,
		"-c", conf,
	)

	cmd.Env = append(os.Environ(),
		"LD_LIBRARY_PATH="+filepath.Join(nginxBase, "libs")+":"+os.Getenv("LD_LIBRARY_PATH"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

////////////////////////////////////////////////////////////////////////////////
// PID HANDLER
////////////////////////////////////////////////////////////////////////////////

func stopByPID(pidFile string) {
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

	if isProcessAlive(pid) {
		proc, _ := os.FindProcess(pid)
		proc.Kill()
	}

	_ = os.Remove(pidFile)
}

func cleanupStalePID(pidFile string) {
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
	if err != nil || !isProcessAlive(pid) {
		fmt.Println("Removing stale PID:", pidFile, "(", pid, ")")
		_ = os.Remove(pidFile)
	}
}

func isProcessAlive(pid int) bool {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = proc.Signal(syscall.Signal(0))
	return err == nil
}

////////////////////////////////////////////////////////////////////////////////
// PORT KILLER
////////////////////////////////////////////////////////////////////////////////

func killPort(port int) {
	out, err := exec.Command("lsof", "-t", "-i", fmt.Sprintf(":%d", port)).Output()
	if err != nil {
		return
	}

	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		pid := strings.TrimSpace(line)
		exec.Command("kill", "-9", pid).Run()
	}
}
