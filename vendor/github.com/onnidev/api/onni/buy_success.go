package onni

import (
	"log"

	"github.com/onnidev/api/types"
)

// BuySuccess TODO: NEEDS COMMENT INFO
func BuySuccess(
	club types.Club,
	party types.Party,
	invoices types.Invoices,
	vouchers []types.Voucher,
	products types.BuyPostList,
	customer types.Customer,
	transactionID string,
) error {
	err := InvoicesBuySuccess(club, party, invoices, products, transactionID)
	if err != nil {
		return err
	}
	log.Println(">>>>>> make vouchers available")
	err = VouchersBuySuccess(vouchers, transactionID)
	if err != nil {
		return err
	}
	err = PartyProductsBuySuccess(products)
	if err != nil {
		return err
	}
	go SendBuySuccessMail(party, club, products, customer)
	return nil

}
