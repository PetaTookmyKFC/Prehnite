package main

import (
	"JsonLanguage/Depot"
	"JsonLanguage/api"
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

	depot, err := Depot.StartDatabase(wd)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// fmt.Println(len(DDepot.Containers))

	// Get Public Folder for templates and files
	PFiles := filepath.Join(wd, "public")
	// Start Server and API
	api.BindAndStartRoutes(":80", PFiles, depot)
}
