package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"evergon/internal/util"
)

type NginxService struct {
	Base    string
	WWWRoot string
}

func NewNginxService(root string, www string) *NginxService {
	return &NginxService{
		Base:    filepath.Join(root, "nginx"),
		WWWRoot: www,
	}
}

func (s *NginxService) Name() string { return "nginx" }

func (s *NginxService) Start() error {
	util.KillPort(8080)
	util.CleanupPID(filepath.Join(s.Base, "logs/nginx.pid"))
	util.PrepareNginxDirs(s.Base)

	// generate nginx.conf from template
	tpl := filepath.Join(s.Base, "conf/nginx.conf.tpl")
	out := filepath.Join(s.Base, "conf/nginx.conf")
	if err := util.GenerateNginxConf(tpl, out, s.WWWRoot); err != nil {
		return err
	}

	nginxBin := filepath.Join(s.Base, "sbin/nginx")
	conf := filepath.Join(s.Base, "conf/nginx.conf")

	fmt.Println("Starting Nginx...")

	cmd := exec.Command(nginxBin,
		"-p", s.Base,
		"-c", conf,
	)

	cmd.Env = append(os.Environ(),
		"LD_LIBRARY_PATH="+filepath.Join(s.Base, "libs")+":"+os.Getenv("LD_LIBRARY_PATH"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *NginxService) Stop() error {
	util.StopPID(filepath.Join(s.Base, "logs/nginx.pid"))
	util.KillPort(8080)
	return nil
}

func (s *NginxService) Status() ServiceStatus {
	pid := util.GetPID(filepath.Join(s.Base, "logs/nginx.pid"))
	return ServiceStatus{
		Running: util.IsAlive(pid),
		PID:     pid,
		Port:    8080,
	}
}
