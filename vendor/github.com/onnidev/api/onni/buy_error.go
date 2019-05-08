package onni

import (
	"log"

	"github.com/onnidev/api/types"
)

// BuyError TODO: NEEDS COMMENT INFO
func BuyError(invoices types.Invoices, vouchers []types.Voucher) error {
	log.Println(">>>>>> make invoices error")
	err := InvoicesBuyError(invoices.All())
	if err != nil {
		log.Println("invoice error", err.Error())
		log.Println("invoice error", err.Error())
		log.Println("invoice error", err.Error())
		log.Println("invoice error", err.Error())
		return err
	}
	log.Println(">>>>>> make vouchers error")
	err = VouchersBuyError(vouchers)
	if err != nil {
		log.Println("voucher error", err.Error())
		log.Println("voucher error", err.Error())
		log.Println("voucher error", err.Error())
		return err
	}
	return nil

}
