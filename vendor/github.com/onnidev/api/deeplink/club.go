package deeplink

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// CLubDeepLInk TODO: NEEDS COMMENT INFO
func CLubDeepLInk(w http.ResponseWriter, r *http.Request) {
	partiesCollection, ok := r.Context().Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	party, err := partiesCollection.GetByID(chi.URLParam(r, "partyID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	buf, err := PartyHTML(party)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	w.WriteHeader(200)
	w.Write(buf.Bytes())
}

// ClubHTML return a html template
func ClubHTML(party types.Club) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("public/clublink.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("ClubDeeplink")
	tmpl, err = tmpl.Parse(string(bb))
	if err != nil {
		return response, err
	}
	err = tmpl.Execute(response, party)
	if err != nil {
		return response, err
	}
	return response, nil
}
