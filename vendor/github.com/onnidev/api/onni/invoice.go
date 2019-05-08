package onni

import (
	"context"

	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
)

// InvoicesFromUser sdkjfn
func InvoicesFromUser(ctx context.Context, id string) ([]types.Invoice, error) {
	invoiceRepo := ctx.Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
	allInvoices, err := invoiceRepo.GetByUserID(id)
	if err != nil {
		return allInvoices, err
	}
	return allInvoices, nil
}

// FutureCustomerInvoices sdkjfn
func FutureCustomerInvoices(ctx context.Context, id string) ([]types.Invoice, error) {
	invoiceRepo := ctx.Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
	allInvoices, err := invoiceRepo.FutureByCustomer(id)
	if err != nil {
		return allInvoices, err
	}
	return allInvoices, nil
}

// Invoice sdkjfn
func Invoice(ctx context.Context, id string) (types.Invoice, error) {
	invoiceRepo := ctx.Value(middlewares.InvoicesRepoKey).(interfaces.InvoicesRepo)
	allInvoices, err := invoiceRepo.GetByID(id)
	if err != nil {
		return allInvoices, err
	}
	return allInvoices, nil
}
