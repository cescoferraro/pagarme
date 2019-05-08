package dates

import (
	"github.com/onnidev/api/report/common"
	"github.com/xuri/excelize"
)

var offset = 2

// Report is sickest
func Report(xlsx *excelize.File, data common.Data) {
	sheet := "Sheet3"
	xlsx.NewSheet(3, sheet)
	xlsx.SetSheetName(sheet, "Datas")
	Headers(xlsx, sheet, data.Vouchers)
	Body(xlsx, sheet, data.Vouchers)
	Charts(xlsx, sheet, data.Vouchers)
}
