package deeplink

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/onnidev/api/shared"
)

// BudDeeplink TODO: NEEDS COMMENT INFO
func BudDeeplink(w http.ResponseWriter, r *http.Request) {
	buf, err := BudBasement()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	w.WriteHeader(200)
	w.Write(buf.Bytes())
}

// BudBasement return a html template
func BudBasement() (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("public/bud.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("BudBasement")
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
