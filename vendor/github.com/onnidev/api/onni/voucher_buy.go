package onni

import (
	"log"

	"github.com/onnidev/api/types"
)

// VouchersBuySuccess TODO: NEEDS COMMENT INFO
func VouchersBuySuccess(vouchers []types.Voucher, transactionID string) error {
	for _, voucher := range vouchers {
		log.Println(voucher)
		err := PatchVoucherBuySuccess(voucher, transactionID)
		if err != nil {
			return err
		}
	}
	return nil
}

// VouchersBuyError TODO: NEEDS COMMENT INFO
func VouchersBuyError(vouchers []types.Voucher) error {
	for _, voucher := range vouchers {
		log.Println(voucher)
		err := PatchVoucherBuyError(voucher)
		if err != nil {
			return err
		}
	}
	return nil
}
