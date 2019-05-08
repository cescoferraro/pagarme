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

// PartyDeeplink TODO: NEEDS COMMENT INFO
func PartyDeeplink(w http.ResponseWriter, r *http.Request) {
	clubsCollection, ok := r.Context().Value(middlewares.CollectionKey).(interfaces.ClubsRepo)
	if !ok {
		err := errors.New("bug")
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	club, err := clubsCollection.GetByID(chi.URLParam(r, "clubID"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	buf, err := ClubHTML(club)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	w.WriteHeader(200)
	w.Write(buf.Bytes())
}

// PartyHTML return a html template
func PartyHTML(party types.Party) (*bytes.Buffer, error) {
	response := new(bytes.Buffer)
	bb, err := shared.Asset("public/partylink.html")
	if err != nil {
		return response, err
	}
	tmpl := template.New("PartyDeeplink")
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
