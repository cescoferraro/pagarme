package summary

import (
	"github.com/onnidev/api/report/common"
	"github.com/xuri/excelize"
)

// Report is sickest
func Report(xlsx *excelize.File, data common.Data) {
	sheet := "Sheet1"
	Headers(sheet, xlsx, data)
	Body(sheet, xlsx, data)
	xlsx.SetSheetName(sheet, "Relatório")
}
