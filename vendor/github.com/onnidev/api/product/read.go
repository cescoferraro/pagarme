package product

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
)

// Read is the shit
func Read(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")
	repo := r.Context().Value(middlewares.ProductsRepoKey).(interfaces.ProductsRepo)
	product, err := repo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, product)
}
