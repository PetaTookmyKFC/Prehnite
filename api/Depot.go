package api

import (
	"JsonLanguage/api/server"
	"encoding/json"
	"fmt"
	"net/http"
)

func DepoAPI() {
	http.HandleFunc("/lst", DepoListApi)
}

func DepoListApi(w http.ResponseWriter, r *http.Request) {
	CM(w, r)
	output := IDepot.ListContainers()

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
	fmt.Println("Responding Data")
}
