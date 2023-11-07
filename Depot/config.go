package Depot

import (
	"encoding/json"
	"fmt"
	"io"
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

	fmt.Println("Config Created!")

}

// func (DD *DBDepot) GenerateConfig() {

//		DD.Containers
//	}
func CreateFromConfig(workingDir string) (bool, *DBDepot) {
	configPath := filepath.Join(workingDir, "config.json")

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("No Config file found")
		return false, nil
	}
	// Try and Open the file
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println("ERROR : - Couldn't open config File")
		os.Exit(1)
	}
	// Auto Close the file
	defer configFile.Close()
	// Read the file
	configData, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println("Error : - Couldn't read config file")
		os.Exit(1)
	}

	var DD DBDepot

	err = json.Unmarshal(configData, &DD)
	if err != nil {
		fmt.Println("Error :- Couldn't parse config File")
		os.Exit(1)
	}
	return true, &DD
}

func (DD *DBDepot) RecalculateConfig() error {
	fmt.Println("Recalculating!")
	// Scan Directory - List all folders!
	folders, err := os.ReadDir(DD.WorkingDir)
	if err != nil {
		return err
	}

	// Clear all containers from the depot
	DD.Containers = make(map[string]DBContainer, 0)

	// Loop threw the folders
	for _, dir := range folders {
		// Loop through items
		if dir.IsDir() {
			// Add the item to the Depot
			DD.AddDBContainer(dir.Name())
		}
	}
	// Save new Config
	DD.SaveConfig()
	return nil
}
