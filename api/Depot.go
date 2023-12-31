package api

import (
	"JsonLanguage/api/server"
	"encoding/json"
	"net/http"
	"path/filepath"
)

func DepoAPI() {
	http.HandleFunc("/api/depot/list", DepoListApi)
	http.HandleFunc("/api/depot/create", DepoCreate)
	http.HandleFunc("/api/depot/recalculate", RecalculateCondfig)
	// http.HandleFunc("/api/depot/info/", DepoInfoApi)
}

type LocationInfo struct {
	Name string `json:"Name"`
}

func DepoCreate(w http.ResponseWriter, r *http.Request) {
	// Check for correct post request
	err := CM(w, r)
	if err != nil {
		return
	}

	var NCLoc LocationInfo

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&NCLoc)
	if err != nil {

		errRes := server.SRE{
			Code:    400,
			Message: err.Error(),
			BJson:   BaseJSON(r),
		}
		errRes.SendHTML()
		return
	}

	if NCLoc.Name == "" {
		s_err := &server.SRE{
			Code:    400,
			Message: "`Name` paramater cannot be empty",
			W:       w,
			BJson:   BaseJSON(r),
		} /////
		s_err.SendHTML()
		return
	}

	IDepot.AddDBContainer(NCLoc.Name)
	w.WriteHeader(200)

	if BaseJSON(r) {
		w.Write([]byte("{message: 'Created Container'}"))
	} else {
		w.Write([]byte("Created Container"))
	}

	IDepot.SaveConfig()

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
				BJson: true,
				Code:  500,
				W:     w,
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

func RecalculateCondfig(w http.ResponseWriter, r *http.Request) {
	err := CM(w, r)
	if err != nil {
		return
	}

	err = IDepot.RecalculateConfig()
	if err != nil {
		htmlError := &server.SRE{
			Code:    500,
			Message: err.Error(),
			BJson:   BaseJSON(r),
		}
		htmlError.SendHTML()
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Recalculated"))

}

// func DepoInfoApi(w http.ResponseWriter, r *http.Request) {
// 	// err := CM(w, r)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	rData := GetWildCardQuery(r.URL.Path, "/api/depot/info/")
// 	// fmt.Printf("GOT CODE : %d", len(rData))
// 	if len(rData) == 0 || rData[0] == "" {

// 		err := server.SRE{
// 			Code:    420,
// 			Message: "/api/depot/info/*MISSING*",
// 			W:       w,
// 		}
// 		err.SendHTML()
// 		return
// 	}

// 	fmt.Println("TestingWildCard : ")

// 	for i, data := range rData {
// 		fmt.Printf(" %d : %s", i, data)
// 	}

// 	container, err := IDepot.GetContainerByName(rData[0])
// 	if err != nil {
// 		err := server.SRE{
// 			Code:    404,
// 			Message: fmt.Sprintf("Container %s not Found", rData[0]),
// 			W:       w,
// 		}
// 		err.SendHTML()
// 		return
// 	}

// 	w.WriteHeader(200)
// 	w.Write([]byte(fmt.Sprintf("Current Conainter contains %d databases", container.Length())))
// }
