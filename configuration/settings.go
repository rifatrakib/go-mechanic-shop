package configuration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

type appConfig struct {
	Port       string                `json:"port"`
	CORSConfig middleware.CORSConfig `json:"cors_config"`
}

func LoadConfig() appConfig {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	config := appConfig{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}

	return config
}
