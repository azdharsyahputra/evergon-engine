package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func hasCap(bin string) bool {
	out, err := exec.Command("getcap", bin).Output()
	if err != nil {
		return false
	}
	return string(out) != ""
}

func (e *Engine) SetupTrust() error {
	nginxBin := filepath.Join(e.BasePath, "nginx", "sbin", "nginx")

	if _, err := os.Stat(nginxBin); err != nil {
		return fmt.Errorf("nginx binary not found")
	}

	// cek dulu, jangan setcap ulang
	if hasCap(nginxBin) {
		fmt.Println("Nginx already trusted")
		return nil
	}

	fmt.Println("[Trust] Granting nginx permission for privileged ports (80/443)")
	fmt.Println("[Trust] This is a one-time setup")

	cmd := exec.Command(
		"sudo",
		"setcap",
		"cap_net_bind_service=+ep",
		nginxBin,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
