package updater

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Version                string             `json:"version"`
	Releases               map[string]Release `json:"releases"`
	ManualQuickDownloadURI string             `json:"quickDownloadURI"`
	ManualFullDownloadURI  string             `json:"fullDownloadURI"`
}

type Release struct {
	Version                 string `json:"version"`
	QuickDownloadURI        string `json:"quickDownloadURI"`
	FullDownloadURI         string `json:"fullDownloadURI"`
	Password                string `json:"password"`
	Name                    string `json:"name"`
	Cabinet                 string `json:"deleteCabinet"`
	MinimumQuickDownloadVer string `json:"minimumQuickDownloadVer"`
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
