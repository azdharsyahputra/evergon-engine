package core

type ProjectConfig struct {
	Name       string `json:"name"`
	PHPVersion string `json:"php_version"`
	Port       int    `json:"port"`
	Root       string `json:"root"`
}
