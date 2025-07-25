package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port           string `json:"port"`
	TaikoDBPath    string `json:"taikoDBPath"`
	AuthDBPath     string `json:"authDBPath"`
	DatatablePath  string `json:"datatablePath"`
	DistPath       string `json:"distPath"`
	LegacyWebUIURL string `json:"legacyWebUIPath"`
	SessionSecret  string `json:"sessionSecret"`
}

var ConfigVars Config

func InitConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ConfigVars)
	if err != nil {
		fmt.Println("Error decoding config:", err)
		return
	}
}
