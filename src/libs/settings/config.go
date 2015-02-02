package settings

import "os"

type Config struct {
	BindPort   string
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	WebRoot    string
}

func (config *Config) Setup() {
	config.BindPort = fetch("PORT", "8080")
	config.DbName = fetch("DBNAME", "mirrorbrain")
	config.DbUser = fetch("DBUSER", "mirrorbrain")
	config.DbPassword = fetch("DBPASSWORD", "pass123")
	config.DbHost = fetch("DBHOST", "127.0.0.1")
	config.WebRoot = fetch("WEBROOT", "/srv/ftp")
}

func fetch(envKey string, defaultValue string) string {
	value := os.Getenv(envKey)
	if value != "" {
		return value
	}
	return defaultValue
}
