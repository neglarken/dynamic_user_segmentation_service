package config

type Config struct {
	Addr        string `yaml:"addr"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
}

func NewConfig(configPath string) *Config {
	return &Config{}
}
