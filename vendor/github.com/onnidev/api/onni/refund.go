package onni

import (
	"context"
	"errors"
	"strings"

	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
)

// ValidateVoucherRefund TODO: NEEDS COMMENT INFO
func ValidateVoucherRefund(ctx context.Context, voucher types.Voucher, user types.UserClub) error {
	if strings.ToLower(user.Profile) == "onni" {
		return nil
	}
	if strings.ToLower(user.Profile) == "admin" {
		if shared.ContainsObjectID(user.Clubs, voucher.ClubID) {
			return nil
		}
		return errors.New("not admin of this specific voucher club")
	}
	return errors.New("sorrry baby")

}
