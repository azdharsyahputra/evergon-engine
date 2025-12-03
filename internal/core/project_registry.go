package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ProjectRegistry struct {
	BasePath string
}

func NewProjectRegistry(base string) *ProjectRegistry {
	return &ProjectRegistry{BasePath: base}
}

// ------------------------------
// PATH HELPERS
// ------------------------------

func (r *ProjectRegistry) projectPath(name string) string {
	return filepath.Join(r.BasePath, "projects", name)
}

func (r *ProjectRegistry) configPath(name string) string {
	return filepath.Join(r.projectPath(name), ".evergon", "config.json")
}

// ------------------------------
// CREATE PROJECT
// ------------------------------

func (r *ProjectRegistry) Create(name string) error {
	root := r.projectPath(name)
	cfgDir := filepath.Join(root, ".evergon")

	// user folders
	if err := os.MkdirAll(filepath.Join(root, "public"), 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfgDir, 0o755); err != nil {
		return err
	}

	cfg := &ProjectConfig{
		Name:       name,
		PHPVersion: "83",
		Port:       10000,
		Root:       "public",
	}

	return r.SaveConfig(name, cfg)
}

// ------------------------------
// LOAD CONFIG
// ------------------------------

func (r *ProjectRegistry) LoadConfig(name string) (*ProjectConfig, error) {
	path := r.configPath(name)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg ProjectConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// ------------------------------
// SAVE CONFIG
// ------------------------------

func (r *ProjectRegistry) SaveConfig(name string, cfg *ProjectConfig) error {
	path := r.configPath(name)

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

// ------------------------------
// UPDATE CONFIG
// ------------------------------

func (r *ProjectRegistry) Update(name string, newCfg *ProjectConfig) error {
	oldCfg, err := r.LoadConfig(name)
	if err != nil {
		return err
	}

	if newCfg.PHPVersion != "" {
		oldCfg.PHPVersion = newCfg.PHPVersion
	}
	if newCfg.Port != 0 {
		oldCfg.Port = newCfg.Port
	}
	if newCfg.Root != "" {
		oldCfg.Root = newCfg.Root
	}

	return r.SaveConfig(name, oldCfg)
}

// ------------------------------
// LIST PROJECTS
// ------------------------------

func (r *ProjectRegistry) List() ([]string, error) {
	dir := filepath.Join(r.BasePath, "projects")

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var list []string
	for _, e := range entries {
		if e.IsDir() {
			list = append(list, e.Name())
		}
	}

	return list, nil
}

// ------------------------------
// LOAD PROJECT ENGINE
// ------------------------------

func (r *ProjectRegistry) Load(name string) (*ProjectEngine, error) {
	if _, err := os.Stat(r.projectPath(name)); err != nil {
		return nil, fmt.Errorf("project not found: %s", name)
	}

	return NewProjectEngine(r.BasePath, name)
}

// ------------------------------
// READ CONFIG (API)
// ------------------------------

func (r *ProjectRegistry) ReadConfig(name string) (*ProjectConfig, error) {
	return r.LoadConfig(name)
}
