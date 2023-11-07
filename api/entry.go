package api

import (
	"JsonLanguage/Depot"
	"JsonLanguage/api/server"
	"net/http"
	"strings"
)

var publicLocation string
var IDepot *Depot.DBDepot
var ApiMethod = "POST"

// Check if the correct method is used
func CM(w http.ResponseWriter, r *http.Request) error {
	if r.Method == ApiMethod {
		return nil
	} else {
		err := &server.SRE{Code: 405, W: w}
		err.SendHTML()
		return err
	}
}
func CMR(w http.ResponseWriter, r *http.Request, override string) error {
	if override == "" {
		err := &server.SRE{Code: 512, W: w}
		err.SendHTML()
		return err
	}
	if r.Method == override {
		return nil
	} else {
		err := &server.SRE{Code: 405, W: w}
		err.SendHTML()
		return err
	}
}

func BindAndStartRoutes(address string, publicDirectory string, Depot *Depot.DBDepot) {
	IDepot = Depot
	publicLocation = publicDirectory
	// List Depots
	DepoAPI()
	server.StartServer(address, publicDirectory)
}

// Check if the request fould be formatted for htmx or just as json
func BaseJSON(r *http.Request) bool {
	// Check if RawQuery constains "BASE"

	if r.URL.Query().Get("B") != "" {
		if strings.ToUpper(r.URL.Query().Get("B")) == "TRUE" {
			return true
		}
	}
	return false
}

// Get the information from the url removing the prefix
func GetWildCardQuery(reqPath string, prefixPath string) (pathArray []string) {
	reqPath = strings.TrimPrefix(reqPath, prefixPath)

	pathArray = strings.Split(reqPath, "/")
	return pathArray
}
