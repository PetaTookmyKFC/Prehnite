package api

import (
	"JsonLanguage/api/server"
	"encoding/json"
	"net/http"
	"path/filepath"
)

func DepoAPI() {
	http.HandleFunc("/api/depot/list", DepoListApi)
	http.HandleFunc("/api/depot/info", DepoInfoApi)
}

func DepoListApi(w http.ResponseWriter, r *http.Request) {
	err := CM(w, r)
	if err != nil {
		return
	}

	output := IDepot.ListContainers()

	// fmt.Println("Running Updated Function!")

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
		// fmt.Println("Responding with Json")
	} else {

		fLoc := filepath.Join("templates", "menu", "DatabaseSelect.html")
		server.SendFile(fLoc, w, output)
		// fmt.Println("Responding with Template")
	}
}

func DepoInfoApi(w http.ResponseWriter, r *http.Request) {
	err := CM(w, r)
	if err != nil {
		return
	}

}
