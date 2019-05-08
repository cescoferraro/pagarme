package clubLead

import (
	"github.com/go-chi/chi"
	"github.com/onnidev/api/middlewares"
)

// Routes is amazing
func Routes(r chi.Router) {
	endpoints := func(n chi.Router) {
		n.Use(middlewares.AttachClubLeadCollection)
		n.Get("/", ListEndpoint)
		n.Get("/{clubLeadID}", ReadEndpoint)
		n.
			With(middlewares.ReadPatchClubLeadRequestFromBody).
			With(middlewares.AttachClubsCollection).
			With(middlewares.AttachGridFSCollection).
			With(middlewares.AttachUserClubCollection).
			With(middlewares.AttachRecipientsCollection).
			Post("/done/{clubLeadID}", Done)
		n.
			With(middlewares.ReadPatchClubLeadRequestFromBody).
			Patch("/{clubLeadID}", Patch)
		n.
			With(middlewares.ReadCreateClubLeadRequestFromBody).
			Post("/", CreateEndpoint)
	}
	r.Route("/clubLead", endpoints)
	r.Route("/clublead", endpoints)

}
