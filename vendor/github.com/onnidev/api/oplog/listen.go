package oplog

import (
	"fmt"
	"log"

	"github.com/onnidev/api/infra"
	"github.com/rwynn/gtm"
)

// Start skjdfn
func Start() {
	store, err := infra.Cloner()
	ctx := gtm.Start(store.Session, nil)
	log.Println("Starting oplog process")
	for {
		select {
		case err = <-ctx.ErrC:
			fmt.Println(err)
		case op := <-ctx.OpC:
			switch op.GetCollection() {
			case "product":
				productOPLOG(op)
				break
			case "voucher":
				voucherOPLOG(op)
				break
			}
		}
	}
}
