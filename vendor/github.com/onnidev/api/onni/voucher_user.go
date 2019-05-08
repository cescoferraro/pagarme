package onni

import (
	"context"
	"errors"
	"time"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// VoucherUseComplete TODO: NEEDS COMMENT INFO
func VoucherUseComplete(ctx context.Context, id string, enforceTime bool, constrain types.VoucherUseConstrain, userClub types.UserClub) (types.CompleteVoucher, error) {
	_, err := VoucherUse(ctx, id, enforceTime, constrain, userClub)
	if err != nil {
		return types.CompleteVoucher{}, err
	}
	vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug")
		return types.CompleteVoucher{}, err
	}
	voucherComplete, err := vouchersCollection.GetByID(id)
	if err != nil {
		return types.CompleteVoucher{}, err
	}
	return voucherComplete, nil
}

// VoucherUse TODO: NEEDS COMMENT INFO
// "read.voucher.error.voucher.not.found"                    : "Esse voucher não foi encontrado ou não tem um código válido. Tente ler novamente.",
// "read.voucher.error.voucher.already.used"                 : "Esse voucher já foi utilizado.",
// "read.voucher.error.voucher.does.not.belong.to.club"      : "Esse voucher não é válido para este clube.",
// "read.voucher.error.voucher.does.not.belong.to.party"     : "Esse voucher não é válido para esta festa.",
// "read.voucher.error.no.current.party"                     : "Esse voucher é válido somente para o dia da festa.",
// "read.voucher.error.voucher.invalid.status"               : "Esse voucher não é válido ou está pendente de pagamento.",
// "read.voucher.error.voucher.invalid.date"                 : "Esse voucher está fora do horário de utilização."
func VoucherUse(ctx context.Context, id string, enforceTime bool, constrains types.VoucherUseConstrain, userClub types.UserClub) (types.Voucher, error) {
	vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug assert")
		return types.Voucher{}, err
	}
	voucher, err := vouchersCollection.GetSimpleByID(id)
	if err != nil {
		return voucher, errors.New("read.voucher.error.voucher.not.found")
	}
	if !constrains.Drink && voucher.Product.Type == "DRINK" {
		err := errors.New("you reader app is not setup to read drinks")
		return voucher, err
	}
	if !constrains.Ticket && voucher.Product.Type == "TICKET" {
		err := errors.New("you reader app is not setup to read tickets")
		return voucher, err
	}
	if userClub.Profile != "ONNI" {
		found := false
		for _, club := range userClub.Clubs {
			if club.Hex() == voucher.ClubID.Hex() {
				found = true
			}
		}
		if !found {
			err := errors.New("read.voucher.error.voucher.does.not.belong.to.club")
			return voucher, err
		}
	}
	err = VoucherCheckStatus(voucher.Status)
	if err != nil {
		return voucher, err
	}
	if enforceTime {
		env := viper.GetString("env")
		if env == "homolog" || env == "production" {
			partyCollection := ctx.Value(middlewares.PartyRepoKey).(interfaces.PartiesRepo)
			party, err := partyCollection.GetByID(voucher.PartyID.Hex())
			if err != nil {
				return voucher, err
			}
			onTime := shared.InTimeSpan(party.StartDate.Time().Add(-time.Hour), party.EndDate.Time(), time.Now())
			if !onTime {
				err := errors.New("read.voucher.error.voucher.invalid.date")
				return voucher, err
			}
		}
		if voucher.Status == "USED" {
			err := errors.New("read.voucher.error.voucher.already.used")
			return voucher, err
		}
	}
	voucher, err = vouchersCollection.UseVoucher(voucher, userClub)
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}

// VoucherCheckStatus TODO: NEEDS COMMENT INFO
func VoucherCheckStatus(status string) error {
	if status == "AVAILABLE" {
		return nil
	}
	if status == "USED" {
		return errors.New("read.voucher.error.voucher.already.used")
	}
	return errors.New("read.voucher.error.voucher.invalid.status")

}
