package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	var config Config

	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(config.Environment)
}
