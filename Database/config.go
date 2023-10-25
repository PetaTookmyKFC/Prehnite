package database

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Save Information about depos into config
func (DD *DBDepot) SaveConfig() {
	configFile := filepath.Join(DD.WorkingDir, "config.json")
	configData, err := json.Marshal(DD)
	if err != nil {
		fmt.Printf("Failed To Compile config file")
		os.Exit(1)
		return
	}

	os.WriteFile(configFile, configData, fs.ModePerm)

	fmt.Print("Created File")

}

// func (DD *DBDepot) GenerateConfig() {

// 	DD.Containers
// }

func (DD *DBDepot) LoadConfig() {

}
