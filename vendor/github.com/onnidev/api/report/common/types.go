package common

import (
	"html/template"
	"os"

	"github.com/onnidev/api/types"
)

// Data sdkjfn
type Data struct {
	Club     types.Club
	Party    types.Party
	Vouchers types.VouchersList
}

// Ok ksjadn
// swagger:response ok
type Ok struct {
	// in: body
	Payload string `json:"body,omitempty"`
}

// DownloadedExcelSheet ksjadn
// swagger:response Binary
type DownloadedExcelSheet struct {
	// type: string
	// format: binary
	// in: body
	Payload os.File `json:"body,omitempty"`
}

// PartyReportHTMLInfo sdfkjn
type PartyReportHTMLInfo struct {
	Name string
	CSS  template.CSS
	Date string
}
