package config

type Config struct {
	HTTPPort int
	FPMPort  int
	APIPort  int
}

func Default() Config {
	return Config{
		HTTPPort: 8080,
		FPMPort:  9099,
		APIPort:  9091,
	}
}
