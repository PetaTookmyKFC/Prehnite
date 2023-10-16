package api

import (
	"JsonLanguage/api/server"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
)

func DepoAPI() {
	http.HandleFunc("/api/lst", DepoListApi)
}

func DepoListApi(w http.ResponseWriter, r *http.Request) {
	CM(w, r)
	output := IDepot.ListContainers()

	fmt.Println("Running Updated Function!")

	if BaseJSON(r) {
		w.WriteHeader(200)
		data, err := json.Marshal(output)
		if err != nil {
			err := &server.SRE{
				Code: 500,
				W:    w,
			}
			err.SendHTML()
			return
		}
		w.Write(data)
		fmt.Println("Responding with Json")
	} else {

		fLoc := filepath.Join("templates", "menu", "DatabaseSelect.html")
		server.SendFile(fLoc, w, output)

		fmt.Println("Responding with Template")
	}
}
