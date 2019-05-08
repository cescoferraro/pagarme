package deeplink

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/onnidev/api/shared"
)

// MainDeeplink TODO: NEEDS COMMENT INFO
func MainDeeplink(w http.ResponseWriter, r *http.Request) {
	buf, err := MainHTML()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	w.WriteHeader(200)
	w.Write(buf.Bytes())
}

// MainHTML return a html template
func MainHTML() (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("public/deeplink.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("MainDeeplink")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	err = tmpl.Execute(response, "hey")
	if err != nil {
		return response, err
	}
	return response, nil
}
