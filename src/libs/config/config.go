package config

import "os"

type Config struct {
	BindPort   string
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
}

func (config *Config) Setup() {
	config.setPort()
	config.DbName = "mirrorbrain"
	config.DbUser = "mirrorbrain"
	config.DbPassword = "mirrorbrain"
	// TODO localhost
	config.DbHost = "192.168.122.136"
}

func (config *Config) setPort() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	config.BindPort = port
}
