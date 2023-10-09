package main

import (
	database "JsonLanguage/Database"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get Database store location
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	depoPath := filepath.Join(wd, "DatabaseDepot")

	DDepot, err := database.StartDatabase(depoPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	DDepot.DatabaseList()

}
