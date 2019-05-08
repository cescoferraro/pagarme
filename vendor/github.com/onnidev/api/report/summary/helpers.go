package summary

import (
	"strconv"

	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
	"github.com/xuri/excelize"
)

func size(all []types.CompleteVoucher) int {
	data := types.VouchersList(all)
	size := len(data.AvailableTickets()) +
		len(data.AvailableFromCategory("BEERS")) +
		len(data.AvailableFromCategory("DOSES")) +
		len(data.AvailableFromCategory("SOFT_DRINKS")) +
		len(data.AvailableFromCategory("SHOTS")) +
		len(data.AvailableFromCategory("VARIOUS")) +
		len(data.AvailableFromCategory("DRINKS"))
	return size + 11
}

func writeRow(data common.Data, xlsx *excelize.File, sheet string, drinks types.VouchersList, row, tipo, kind string) {
	drinks = drinks.FilterByNotStatus("ERROR")
	xlsx.SetCellValue(sheet, "A"+row, tipo)
	xlsx.SetCellValue(sheet, "B"+row, drinks.FilterByNotStatus("TRANSFERED").Size())
	xlsx.SetCellValue(sheet, "D"+row, drinks.
		FilterByTypes("NORMAL", "TRANSFERED").
		FilterByStatuses("CANCELED", "USED", "AVAILABLE").Size())
	xlsx.SetCellValue(sheet, "E"+row, drinks.FilterByType("EXTERNAL_BUY").Size())
	xlsx.SetCellValue(sheet, "F"+row, drinks.FilterByType("PROMOTION").Size())
	xlsx.SetCellValue(sheet, "G"+row, drinks.FilterByType("FREE").Size())
	xlsx.SetCellValue(sheet, "H"+row, drinks.FilterByType("ANNIVERSARY").Size())
	xlsx.SetCellValue(sheet, "J"+row, drinks.FilterByStatus("AVAILABLE").Size())
	xlsx.SetCellValue(sheet, "K"+row, drinks.FilterByStatus("USED").Size())
	xlsx.SetCellValue(sheet, "L"+row, drinks.FilterByStatus("CANCELED").Size())
	xlsx.SetCellValue(sheet, "N"+row, drinks.Accountable().Sum())
	var liquidFloat float64
	if kind == "DRINKS" {
		liquidFloat = drinks.Accountable().LiquidSumDrinks(data.Club.PercentDrink)
	}
	if kind == "INGRESSO" {
		liquidFloat = drinks.Accountable().LiquidSumTickets(data.Club.PercentTicket, data.Party.AssumeServiceFee)
	}
	xlsx.SetCellValue(sheet, "O"+row, liquidFloat)
}

// Headers sdkjfn
func Headers(sheet string, xlsx *excelize.File, data common.Data) {
	size := size(data.Vouchers)
	styling(data.Party, sheet, size, xlsx)
	resumos(sheet, xlsx, data)
}

// Body sdfkjn
func Body(sheet string, xlsx *excelize.File, data common.Data) {
	last := ticketsReport(data, xlsx, sheet)
	last = genericReport(data, xlsx, sheet, "BEERS", "CERVEJAS", last)
	last = genericReport(data, xlsx, sheet, "DOSES", "DOSES", last)
	last = genericReport(data, xlsx, sheet, "SHOTS", "SHOTS", last)
	last = genericReport(data, xlsx, sheet, "DRINKS", "DRINKS", last)
	last = genericReport(data, xlsx, sheet, "VARIOUS", "DIVERSOS", last)
	last = genericReport(data, xlsx, sheet, "SOFT_DRINKS", "REFRIGERANTES", last)
	last = genericReport(data, xlsx, sheet, "BOTTLES", "GARRAFAS", last)
	genericReport(data, xlsx, sheet, "COMBOS", "COMBOS", last)
}

func styling(party types.Party, sheet string, size int, xlsx *excelize.File) {
	xlsx.SetCellValue(sheet, "A2", party.Name)
	labels(party, sheet, xlsx)
	totalHeightHeader(size, sheet, xlsx)
	black, _ := common.BlackTable(xlsx)
	xlsx.SetCellStyle(sheet, "A1", "O2", black)
	xlsx.SetCellStyle(sheet, "A2", "A2", black)
	xlsx.SetCellStyle(sheet, "A4", "O4", black)
	xlsx.SetCellStyle(sheet, "A6", "O6", black)
	xlsx.SetCellStyle(sheet, "A9", "O9", black)
	xlsx.SetCellStyle(sheet, "A18", "O18", black)
	// xlsx.SetCellStyle(sheet, "A16", "O16", black)
	xlsx.SetCellStyle(sheet, "C1", "C"+strconv.Itoa(offset+size), black)
	xlsx.SetCellStyle(sheet, "I1", "I"+strconv.Itoa(offset+size), black)
	xlsx.SetCellStyle(sheet, "M1", "M"+strconv.Itoa(offset+size), black)
	str, _ := common.ReportStyle(xlsx)
	xlsx.SetCellStyle(sheet, "B5", "B5", str)
	xlsx.SetCellStyle(sheet, "D5", "H5", str)
	xlsx.SetCellStyle(sheet, "J5", "L5", str)
	xlsx.SetCellStyle(sheet, "A"+strconv.Itoa(offset+1), "B"+strconv.Itoa(offset+size+1), str)
	xlsx.SetCellStyle(sheet, "D"+strconv.Itoa(offset+1), "H"+strconv.Itoa(offset+size+1), str)
	xlsx.SetCellStyle(sheet, "J"+strconv.Itoa(offset+1), "L"+strconv.Itoa(offset+size+1), str)
	curency, _ := common.CurrencyTable(xlsx)
	xlsx.SetCellStyle(sheet, "N5", "O5", curency)
	xlsx.SetCellStyle(sheet, "N"+strconv.Itoa(offset+1), "O"+strconv.Itoa(offset+size+1), curency)
	widthHeader(sheet, xlsx)
}
func labels(party types.Party, sheet string, xlsx *excelize.File) {
	str, _ := common.HeaderStyle(xlsx)
	xlsx.SetCellStyle(sheet, "A3", "O3", str)
	xlsx.SetCellStyle(sheet, "A5", "A5", str)
	xlsx.SetCellValue(sheet, "A5", "TOTAL")
	xlsx.SetCellValue(sheet, "A3", "PRODUTO")
	xlsx.SetCellValue(sheet, "B3", "TOTAL")
	xlsx.SetCellValue(sheet, "D3", "VENDA")
	xlsx.SetCellValue(sheet, "E3", "VENDA EXTERNA")
	xlsx.SetCellValue(sheet, "F3", "PROMOÇÃO")
	xlsx.SetCellValue(sheet, "G3", "FREE")
	xlsx.SetCellValue(sheet, "H3", "ANIVERSÁRIO")
	xlsx.SetCellValue(sheet, "J3", "DISPONÏVEL")
	xlsx.SetCellValue(sheet, "K3", "USADO")
	xlsx.SetCellValue(sheet, "L3", "CANCELADO")
	xlsx.SetCellValue(sheet, "N3", "RECEITA BRUTA")
	xlsx.SetCellValue(sheet, "O3", "RECEITA LÍQUIDA")

}

func widthHeader(sheet string, xlsx *excelize.File) {
	xlsx.SetColWidth(sheet, "A", "A", 35)
	xlsx.SetColWidth(sheet, "B", "B", 10)
	xlsx.SetColWidth(sheet, "C", "C", 1)
	xlsx.SetColWidth(sheet, "D", "H", 10)
	xlsx.SetColWidth(sheet, "I", "I", 1)
	xlsx.SetColWidth(sheet, "J", "L", 10)
	xlsx.SetColWidth(sheet, "M", "M", 1)
	xlsx.SetColWidth(sheet, "N", "O", 15)
}

func totalHeader(sheet string, size int, xlsx *excelize.File) {
	cols := []string{"B", "D", "E", "F", "G", "H", "J", "K", "L", "N", "O"}
	for _, col := range cols {
		xlsx.SetCellFormula(
			sheet,
			col+strconv.Itoa(offset-1),
			"SUM("+col+strconv.Itoa(offset+1)+":"+col+strconv.Itoa(offset+size)+")")
	}
}

func totalHeightHeader(size int, sheet string, xlsx *excelize.File) {
	for i := 0; i <= offset+size; i++ {
		xlsx.SetRowHeight(sheet, i, 20)
	}
	xlsx.SetRowHeight(sheet, 2, 45)
}

func resumos(sheet string, xlsx *excelize.File, data common.Data) {
	resumo(xlsx, data, sheet, "5", "TOTAL", "TOTAL", "TOTAL")
	resumo(xlsx, data, sheet, "7", "TICKET", "INGRESSO", "SUBTOTAL INGRESSO")
	resumo(xlsx, data, sheet, "8", "CONSUMO", "CONSUMO", "SUBTOTAL DRINKS")

	resumo(xlsx, data, sheet, "10", "BEERS", "DRINKS", "CERVEJAS")
	resumo(xlsx, data, sheet, "12", "DOSES", "DRINKS", "DOSES")
	resumo(xlsx, data, sheet, "11", "SHOTS", "DRINKS", "SHOTS")
	resumo(xlsx, data, sheet, "13", "DRINKS", "DRINKS", "DRINKS")
	resumo(xlsx, data, sheet, "14", "VARIOUS", "DRINKS", "DIVERSOS")
	resumo(xlsx, data, sheet, "15", "SOFT_DRINKS", "DRINKS", "REFRIGERANTES")
	resumo(xlsx, data, sheet, "16", "BOTTLES", "DRINKS", "GARRAFAS")
	resumo(xlsx, data, sheet, "17", "COMBOS", "DRINKS", "COMBOS")
}

// WriteSubHeader whatever
func WriteSubHeader(xlsx *excelize.File, sheet, name, target string) {
	xlsx.SetCellValue(sheet, "A"+target, name)
	black, _ := common.SubHeaderStyle(xlsx)
	xlsx.SetCellStyle(sheet, "A"+target, "O"+target, black)
	xlsx.MergeCell(sheet, "A"+target, "O"+target)
}

const offset = 18
