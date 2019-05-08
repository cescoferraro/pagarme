package dates

import (
	"strconv"

	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
	"github.com/xuri/excelize"
)

// Headers sdkjfn
func Headers(xlsx *excelize.File, sheet string, all []types.CompleteVoucher) {
	list := types.VouchersList(all)
	hey := list.ConsumedDates()
	xlsx.SetColWidth(sheet, "A", "A", 20)
	xlsx.SetColWidth(sheet, "B", "G", 15)
	xlsx.SetCellValue(sheet, "A1", "Datas")
	xlsx.SetCellValue(sheet, "B1", "Ticket Qtd")
	xlsx.SetCellValue(sheet, "C1", "Ticket R$")
	xlsx.SetCellValue(sheet, "D1", "Drinks Qtd")
	xlsx.SetCellValue(sheet, "E1", "Drinks R$")
	xlsx.SetCellValue(sheet, "F1", "Total Qtd")
	xlsx.SetCellValue(sheet, "G1", "Total R$")
	str, _ := common.ReportStyle(xlsx)
	header, _ := common.HeaderStyle(xlsx)
	currency, _ := common.CurrencyTable(xlsx)
	black, _ := common.BlackTable(xlsx)
	last := strconv.Itoa(offset + len(hey))
	first := strconv.Itoa(offset)
	xlsx.SetCellStyle(sheet, "A1", "G1", black)
	xlsx.SetCellStyle(sheet, "A2", "A"+last, header)
	xlsx.SetCellStyle(sheet, "B"+first, "B"+last, str)
	xlsx.SetCellStyle(sheet, "D"+first, "D"+last, str)
	xlsx.SetCellStyle(sheet, "F"+first, "F"+last, str)
	xlsx.SetCellStyle(sheet, "C"+first, "C"+last, currency)
	xlsx.SetCellStyle(sheet, "E"+first, "E"+last, currency)
	xlsx.SetCellStyle(sheet, "G"+first, "G"+last, currency)
	for i := 0; i < offset+1+len(hey)+1; i++ {
		xlsx.SetRowHeight(sheet, i, 25)
	}
}
