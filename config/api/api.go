package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigApi struct {
	Api ApiData
}

type ApiData struct {
	Port string
}

func GetConfig() (port string) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	var configApi ConfigApi

	err := decoder.Decode(&configApi)
	if err != nil {
		fmt.Println("error:", err)
	}

	port = configApi.Api.Port
	return
}
