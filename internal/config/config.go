package config

type Config struct {
	Completion string
	NoCopy     bool
	Prefix     string
}

func New() *Config {
	return &Config{}
}
