package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Config เก็บ configuration ของแอปพลิเคชัน
type Config struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}

var configInstance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{
			AppName: "My REST API",
			Port:    8080,
		}
	})
	return configInstance
}

func main() {
	config := GetConfig()
	config2 := GetConfig()

	fmt.Println(config == config2) // true

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
