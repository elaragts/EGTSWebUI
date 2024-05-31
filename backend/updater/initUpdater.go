package updater

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Version  string             `json:"version"`
	Releases map[string]Release `json:"releases"`
}

type Release struct {
	Version  string `json:"version"`
	URI      string `json:"uri"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Cabinet  string `json:"deleteCabinet"`
}

var UpdaterVars Config

func InitConfig() {
	file, err := os.Open("updater/config.json")
	if err != nil {
		fmt.Println("Error opening updater config file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&UpdaterVars)
	if err != nil {
		fmt.Println("Error decoding updater config:", err)
		return
	}
}
