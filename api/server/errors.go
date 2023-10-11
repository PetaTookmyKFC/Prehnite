package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var ApiMethod = "GET"

type SRE struct {
	Code int
	W    http.ResponseWriter
}

func (se *SRE) SendHTML() error {
	// Send Error file placed in `Public/assets/codePages/` depending on the error
	se.W.WriteHeader(se.Code)
	fname := fmt.Sprintf("%d.html", se.Code)
	fPath := filepath.Join(publicDir, "assets", "codePages", fname)

	tmp, err := template.ParseFiles(fPath)
	if err != nil {
		se.W.WriteHeader(500)
		se.W.Write([]byte(err.Error()))
		return err
	}
	// Create Template
	err = tmp.Execute(se.W, nil)
	if err != nil {
		return err
	}

	return nil
}
func (se *SRE) Error() string {

	// Send Correct responce code
	switch se.Code {
	case 405:
		{
			return "Method Not Allowed"
		}
	case 404:
		{
			return "File Not Found"
		}
	case 500:
		{
			return "Internal Server Error"
		}
	}
	return "Unknown Error Occoured"
}
