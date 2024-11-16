package config

type Config struct {
	NoCopy bool
	Prefix string
}

func New() *Config {
	return &Config{}
}
