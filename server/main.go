package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var publicDir string

func sendFile(file string, w http.ResponseWriter, data interface{}) error {
	fPath := filepath.Join(publicDir, file)
	// Check file exist
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		// Handle Not Exist
		w.WriteHeader(404)
		w.Write([]byte("File Not Found! "))
		return err
	} else if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	// Create template from file
	tmp, err := template.ParseFiles(fPath)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return err
	}

	err = tmp.Execute(w, data)
	return err
}

func StartServer(address string, publicDirectory string) {
	publicDir = publicDirectory
	// Main Web Container
	http.HandleFunc("/", MainPage)

	/*
		Start Asset file serving
			1. Get Path to asset folder
			2. Create a Fileserver for the assets
			3. Remove the prefix used for the assets route
			4. Server static Assets
	*/
	statPath := filepath.Join(publicDir, "assets")
	staticAssets := http.FileServer(http.Dir(statPath))
	staticHandler := http.StripPrefix("/a/", staticAssets)
	http.Handle("/a/", staticHandler)

	// Host the Server
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	sendFile("MAIN.html", w, nil)
}
