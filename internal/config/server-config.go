package config

type ServerConfig struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Templates string `yaml:"templates"`
}
