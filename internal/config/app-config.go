package config

type AppConfig struct {
	DBConfig     *DBConfig     `yaml:"dbconnection"`
	ServerConfig *ServerConfig `yaml:"server"`
}
