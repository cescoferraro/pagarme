package list

import (
	"strconv"

	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
	"github.com/xuri/excelize"
)

// Report is sick
func Report(xlsx *excelize.File, data common.Data) {
	xlsx.NewSheet(2, "Sheet2")
	// VouchersSetPanes(xlsx)
	writeHeaders(xlsx)
	setTableStyles("Sheet2", xlsx, (data.Vouchers))
	errorLess := data.Vouchers.ExcludeErrors()
	for index, ele := range errorLess {
		writeRow("Sheet2", xlsx, index, ele)
	}
	xlsx.SetSheetName("Sheet2", "Vouchers")
	xlsx.SetActiveSheet(2)
}

func writeHeaders(xlsx *excelize.File) {
	style, _ := common.HeaderStyle(xlsx)
	xlsx.SetCellStyle("Sheet2", "A1", "K2", style)
	xlsx.MergeCell("Sheet2", "A1", "A2")
	xlsx.MergeCell("Sheet2", "B1", "B2")
	xlsx.MergeCell("Sheet2", "C1", "C2")
	xlsx.MergeCell("Sheet2", "D1", "D2")
	xlsx.MergeCell("Sheet2", "E1", "E2")
	xlsx.MergeCell("Sheet2", "F1", "F2")
	xlsx.MergeCell("Sheet2", "G1", "G2")
	xlsx.MergeCell("Sheet2", "H1", "H2")
	xlsx.MergeCell("Sheet2", "I1", "I2")
	xlsx.MergeCell("Sheet2", "J1", "J2")
	xlsx.MergeCell("Sheet2", "K1", "K2")
	xlsx.SetRowHeight("Sheet2", 0, 25)
	xlsx.SetRowHeight("Sheet2", 1, 25)
	xlsx.SetColWidth("Sheet2", "A", "A", 23)
	xlsx.SetColWidth("Sheet2", "B", "B", 15)
	xlsx.SetColWidth("Sheet2", "C", "C", 38)
	xlsx.SetColWidth("Sheet2", "D", "D", 42)
	xlsx.SetColWidth("Sheet2", "E", "E", 35)
	xlsx.SetColWidth("Sheet2", "F", "F", 50)
	xlsx.SetColWidth("Sheet2", "G", "G", 28)
	xlsx.SetColWidth("Sheet2", "H", "H", 25)
	xlsx.SetColWidth("Sheet2", "I", "I", 23)
	xlsx.SetColWidth("Sheet2", "J", "J", 30)
	xlsx.SetColWidth("Sheet2", "K", "K", 30)
	xlsx.SetCellValue("Sheet2", "A2", "DATA")
	xlsx.SetCellValue("Sheet2", "B2", "STATUS")
	xlsx.SetCellValue("Sheet2", "C2", "CLIENTE")
	xlsx.SetCellValue("Sheet2", "D2", "E-MAIL")
	xlsx.SetCellValue("Sheet2", "E2", "TIPO DE PRODUTO")
	xlsx.SetCellValue("Sheet2", "F2", "PRODUTO")
	xlsx.SetCellValue("Sheet2", "G2", "VALOR")
	xlsx.SetCellValue("Sheet2", "H2", "TIPO DE VOUCHER")
	xlsx.SetCellValue("Sheet2", "I2", "RESPONSÁVEL")
	xlsx.SetCellValue("Sheet2", "J2", "DATA DE USO")
	xlsx.SetCellValue("Sheet2", "K2", "OPERADOR")
}

func setTableStyles(sheet string, xlsx *excelize.File, vouchers []types.CompleteVoucher) {
	size := len(vouchers)
	style, _ := common.Table(xlsx)
	dateStyle, _ := common.DateTable(xlsx)
	pricesStyle, _ := common.CurrencyTable(xlsx)
	xlsx.SetCellStyle(sheet, "A3", "A"+strconv.Itoa(size+2), dateStyle)
	xlsx.SetCellStyle(sheet, "B3", "F"+strconv.Itoa(size+2), style)
	xlsx.SetCellStyle(sheet, "G3", "G"+strconv.Itoa(size+2), pricesStyle)
	xlsx.SetCellStyle(sheet, "H3", "I"+strconv.Itoa(size+2), style)
	xlsx.SetCellStyle(sheet, "J3", "H"+strconv.Itoa(size+2), dateStyle)
	xlsx.SetCellStyle(sheet, "K3", "K"+strconv.Itoa(size+2), style)
	// xlsx.AddTable("Sheet2", "A2", "K"+strconv.Itoa(len(vouchers)), ``)
}

func writeRow(sheet string, xlsx *excelize.File, index int, voucher types.CompleteVoucher) {
	xlsx.SetRowHeight(sheet, index+2, 25)
	xlsx.SetCellValue(sheet, "A"+add2(index), voucher.CreationDate.Time())
	xlsx.SetCellValue(sheet, "B"+add2(index), transformStatus(voucher.Status))
	xlsx.SetCellValue(sheet, "C"+add2(index), voucher.CustomerName)
	if voucher.Customer != nil {
		xlsx.SetCellValue(sheet, "D"+add2(index), voucher.Customer.Mail)
	}
	xlsx.SetCellValue(sheet, "E"+add2(index), transformVoucherType(voucher.Product.Type))
	xlsx.SetCellValue(sheet, "F"+add2(index), voucher.Product.Name)
	xlsx.SetCellValue(sheet, "G"+add2(index), voucher.Price.Value)
	xlsx.SetCellValue(sheet, "H"+add2(index), transformProductType(voucher.Type))
	if voucher.Responsable != nil {
		xlsx.SetCellValue(sheet, "I"+add2(index), voucher.Responsable.Name)
	}
	if voucher.VoucherUseDate != nil {
		xlsx.SetCellValue(sheet, "J"+add2(index), voucher.VoucherUseDate.Time())
	}
	xlsx.SetCellValue(sheet, "K"+add2(index), voucher.VoucherUseUserClubName)
}

func add2(index int) string {
	return strconv.Itoa(index + 3)
}

func transformVoucherType(typ string) string {
	switch typ {
	case "TICKET":
		return "INGRESSO"
	case "DRINK":
		return "CONSUMO"
	}
	return "ERRO EXCEL"
}

func transformProductType(typ string) string {
	switch typ {
	case "NORMAL":
		return "VENDA"
	case "PROMOTION":
		return "PROMOÇÃO"
	case "ANNIVERSARY":
		return "FREE ANIVERSÁRIO"
	case "FREE":
		return "FREE CORTESIA"
	case "TRANSFERED":
		return "TRANSFERIDO"
	case "EXTERNAL_BUY":
		return "VENDA EXTERNA"
	}
	return "ERRO EXCEL"
}

func transformStatus(status string) string {
	switch status {
	case "CANCELED":
		return "CANCELADO"
	case "AVAILABLE":
		return "DISPONÍVEL"
	case "PROCESSING":
		return "PROCESSANDO"
	case "PENDING":
		return "PENDENTE"
	case "USED":
		return "USADO"
	case "TRANSFERED":
		return "TRANSFERIDO"
	case "ERROR":
		return "ERRO NA COMPRA"
	}
	return "ERRO EXCEL"
}
