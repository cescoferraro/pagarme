package voucher

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Transfer is commented
// swagger:route POST /vouchers/transfer/{voucherId} backoffice transferVoucher
//
// Transferencia de voucher entre usuários:
// Vouchers que podem ser transferidos: type = Compra ou Promocional e status = Disponível
// Quando um voucher é transferido o voucher original assume o STATUS transferido e um novo voucher com TYPE transferido é criado.
// O novo voucher é praticamente um clone do voucher original, com exceção de alguns campos:
// Campos alterados:
// "creationDate"
// "updateDate"
// "customerId"
// "customerName"
// "status": avaliable
// "type": transfered
//   Novo campo:
//     "transferedFrom": id do voucher original
// Leitura de voucher transferido: vouchers com STATUS transferido não podem ser validados.
// Estorno de voucher transferido: o estorno fica habilitado para o novo voucher, mas o serviço chama o id do voucher original para realizar o estorno utilizando o id informado no campo "transferedFrom".
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: https
//
//     Security:
//       JWT_TOKEN:
//
//     Responses:
//       200: voucherType
func Transfer(w http.ResponseWriter, r *http.Request) {
	// clubUser := r.Context().Value(middlewares.UserClubKey).(types.UserClub)
	vouchersCollection := r.Context().Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	id := chi.URLParam(r, "id")
	voucher, err := vouchersCollection.GetSimpleByID(id)
	if err != nil {
		render.Status(r, http.StatusPreconditionFailed)
		render.JSON(w, r, http.StatusText(http.StatusPreconditionFailed))
		return
	}
	voucherLogger("Voucher Requested", voucher)
	if voucher.Status == "AVAILABLE" && (voucher.Type == "NORMAL" || voucher.Type == "PROMOTION") {
		customerReceiving := r.Context().Value(middlewares.TranferableCustomerKey).(types.Customer)
		newVoucher := voucher.ToBeTransfed(customerReceiving)
		voucherLogger("Voucher To be inserted", voucher)
		err := vouchersCollection.Insert(newVoucher)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, http.StatusText(http.StatusBadRequest))
			return
		}
		canceledVoucher, err := vouchersCollection.GetByIDAndTransfer(id)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, http.StatusText(http.StatusBadRequest))
			return
		}
		voucherLogger("Voucher canceled", voucher)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, canceledVoucher)

	} else {
		log.Println(222)
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
	}
}

func voucherLogger(label string, voucher types.Voucher) {
	if viper.GetBool("verbose") {
		j, _ := json.MarshalIndent(voucher, "", "    ")
		log.Println(label)
		log.Println(string(j))
	}
}
