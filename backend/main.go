package main

import (
	"encoding/json"
	"fmt"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/cmd/api"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"os"
)

type Config struct {
	Port        string `json:"port"`
	TaikoDBPath string `json:"taikoDBPath"`
	AuthDBPath  string `json:"authDBPath"`
	DistPath    string `json:"distPath"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config:", err)
		return
	}
	database.InitDBs(config.TaikoDBPath, config.AuthDBPath)
	api.Run(config.Port, config.DistPath)
}
