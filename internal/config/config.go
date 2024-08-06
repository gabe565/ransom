package config

type Config struct {
	NoCopy bool
}

func New() *Config {
	return &Config{}
}
