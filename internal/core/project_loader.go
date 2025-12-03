package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Load single project config
func LoadProjectConfig(base, name string) (*ProjectConfig, error) {
	path := filepath.Join(base, "projects", name, ".evergon", "config.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("project config not found: %w", err)
	}

	var cfg ProjectConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid project config: %w", err)
	}

	return &cfg, nil
}
func SaveProjectConfig(base, name string, cfg *ProjectConfig) error {
	path := filepath.Join(base, "projects", name, ".evergon", "config.json")

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}
