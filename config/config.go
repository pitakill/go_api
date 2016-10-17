package config

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

type Config struct {
	Environment Environment
}

type Environment struct {
	Devel Data
}

type Data struct {
	Database,
	Port,
	User,
	Password string
	System OS
}

type OS struct {
	Linux, Darwin Host
}

type Host struct {
	IP string
}

type DataConfig struct {
	User,
	Password,
	Database,
	Port,
	IP string
}

func NewConfig(data Config) DataConfig {
	config := DataConfig{}
	environment := os.Getenv("GO_API_ENV")

	if environment == "devel" || environment == "" {
		config.Port = data.Environment.Devel.Port
		config.User = data.Environment.Devel.User
		config.Database = data.Environment.Devel.Database
		config.Password = data.Environment.Devel.Password

		switch os := runtime.GOOS; os {
		case "linux":
			config.IP = data.Environment.Devel.System.Linux.IP
		case "darwin":
			config.IP = data.Environment.Devel.System.Darwin.IP
		}
	}

	return config
}

func GetConfig() DataConfig {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	var config Config

	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return NewConfig(config)
}
