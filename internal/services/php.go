package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"evergon/internal/util"
)

type PHPService struct {
	Base string
}

func NewPHPService(root string) *PHPService {
	return &PHPService{
		Base: filepath.Join(root, "php"),
	}
}

func (s *PHPService) Name() string { return "php-fpm" }

func (s *PHPService) Start() error {
	util.KillPort(9099)
	util.CleanupPID(filepath.Join(s.Base, "logs/php-fpm.pid"))
	util.PreparePHPDirs(s.Base)

	fpmBin := filepath.Join(s.Base, "sbin/php-fpm")
	conf := filepath.Join(s.Base, "etc/php-fpm.conf")
	ini := filepath.Join(s.Base, "etc/php.ini")

	fmt.Println("Starting PHP-FPM...")

	cmd := exec.Command(fpmBin,
		"-p", s.Base,
		"-y", conf,
		"-c", ini,
		"--daemonize",
	)

	cmd.Env = append(os.Environ(),
		"LD_LIBRARY_PATH="+filepath.Join(s.Base, "libs")+":"+os.Getenv("LD_LIBRARY_PATH"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *PHPService) Stop() error {
	util.StopPID(filepath.Join(s.Base, "logs/php-fpm.pid"))
	util.KillPort(9099)
	return nil
}

func (s *PHPService) Status() ServiceStatus {
	pid := util.GetPID(filepath.Join(s.Base, "logs/php-fpm.pid"))
	return ServiceStatus{
		Running: util.IsAlive(pid),
		PID:     pid,
		Port:    9099,
	}
}
