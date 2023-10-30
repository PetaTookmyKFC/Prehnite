package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var ApiMethod = "GET"

type SRE struct {
	Code    int
	Message string
	BJson   bool
	W       http.ResponseWriter
}

func BaseJSON(r *http.Request) bool {
	// Check if RawQuery constains "BASE"

	if r.URL.Query().Get("B") != "" {
		if strings.ToUpper(r.URL.Query().Get("B")) == "TRUE" {
			return true
		}
	}
	return false
}

func (se *SRE) SendHTML() error {
	// Send Error file placed in `Public/assets/codePages/` depending on the error
	se.W.WriteHeader(se.Code)
	fname := fmt.Sprintf("%d.html", se.Code)

	if se.BJson {
		se.W.WriteHeader(se.Code)
		se.W.Write([]byte("{Message: '" + se.Error() + " - " + se.Message + "'}"))
		return nil
	}

	fPath := filepath.Join(publicDir, "assets", "codePages", fname)

	tmp, err := template.ParseFiles(fPath)
	if err != nil {
		se.W.WriteHeader(500)
		se.W.Write([]byte(err.Error()))
		return err
	}
	// Create Template
	err = tmp.Execute(se.W, se.Message)
	if err != nil {
		return err
	}

	return nil
}
func (se *SRE) Error() string {

	// Send Correct responce code
	switch se.Code {
	case 400:
		{
			return "Bad Request"
		}
	case 405:
		{
			return "Method Not Allowed"
		}
	case 404:
		{
			return "File Not Found"
		}
	case 420:
		{
			return "Missing URL wildcard"
		}
	case 500:
		{
			return "Internal Server Error"
		}
	}
	return "Unknown Error Occoured"
}
