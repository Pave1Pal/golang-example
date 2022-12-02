package config

type DBConfig struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	DBName            string `yaml:"dbname"`
	Password          string `yaml:"password"`
	UserName          string `yaml:"username"`
	SslMode           string `yaml:"sslmode"`
	DriverName        string `yaml:"driver-name"`
	MigrationFilesURL string `yaml:"migration-files-url"`
}
