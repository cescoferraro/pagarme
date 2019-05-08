package oplog

import (
	"net/http"

	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
)

// InfoOP skdjfn
type InfoOP struct {
	Count       int
	LastProduct types.Product
	LastVoucher types.Voucher
}

// OPLOGGLOBAL sdkjfn
var OPLOGGLOBAL = InfoOP{
	Count:       0,
	LastProduct: types.Product{},
	LastVoucher: types.Voucher{},
}

// Info TODO: NEEDS COMMENT INFO
func Info(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, OPLOGGLOBAL)
}
