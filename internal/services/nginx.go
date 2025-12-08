package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	// 1. Bersihin port + PID lama
	util.KillPort(8080)
	util.CleanupPID(filepath.Join(s.Base, "logs/nginx.pid"))
	util.PrepareNginxDirs(s.Base)

	// 2. Generate nginx.conf dinamis (tanpa template eksternal)
	confFile := filepath.Join(s.Base, "conf/nginx.conf")
	if err := s.generateConfig(confFile); err != nil {
		return err
	}

	nginxBin := filepath.Join(s.Base, "sbin/nginx")
	conf := confFile

	fmt.Println("Starting Nginx global (www mode)...")

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

// -------------------------------------------
// CONFIG GENERATOR
// -------------------------------------------

func (s *NginxService) generateConfig(outPath string) error {
	mimeTypes := filepath.Join(s.Base, "conf/mime.types")
	fastcgiConf := filepath.Join(s.Base, "conf/fastcgi.conf")
	logDir := filepath.Join(s.Base, "logs")

	serverBlocks, err := s.buildServerBlocks(fastcgiConf)
	if err != nil {
		return err
	}

	if serverBlocks == "" {
		// fallback minimal kalau www kosong
		serverBlocks = fmt.Sprintf(`
server {
    listen 8080 default_server;
    server_name localhost;
    root %s;
    index index.php index.html;

    location / {
        try_files $uri $uri/ =404;
    }
}
`, filepath.Join(s.Base, "html"))
	}

	conf := fmt.Sprintf(`
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       %s;
    default_type  application/octet-stream;

    access_log  %s/access.log;
    error_log   %s/error.log;

%s
}
`, mimeTypes, logDir, logDir, serverBlocks)

	return os.WriteFile(outPath, []byte(conf), 0644)
}

func (s *NginxService) buildServerBlocks(fastcgiConf string) (string, error) {
	entries, err := os.ReadDir(s.WWWRoot)
	if err != nil {
		// kalau folder www belum ada / belum kepakai, anggap kosong
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	var servers []string
	first := true

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		name := e.Name()
		siteRoot := filepath.Join(s.WWWRoot, name)

		publicRoot := filepath.Join(siteRoot, "public")
		if st, err := os.Stat(publicRoot); err == nil && st.IsDir() {
			siteRoot = publicRoot
		}

		listen := "8080"
		if first {
			listen = "8080 default_server"
			first = false
		}

		server := fmt.Sprintf(`
server {
    listen %s;
    server_name %s.local;

    root %s;
    index index.php index.html;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location ~ \.php$ {
        include %s;
        fastcgi_pass 127.0.0.1:9099;
        fastcgi_param SCRIPT_FILENAME %s$fastcgi_script_name;
    }
}
`, listen, name, siteRoot, fastcgiConf, siteRoot)

		servers = append(servers, server)
	}

	return strings.Join(servers, "\n"), nil
}
