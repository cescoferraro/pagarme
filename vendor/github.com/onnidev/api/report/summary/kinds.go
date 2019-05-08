package summary

import (
	"strconv"

	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
	"github.com/xuri/excelize"
)

func ticketsReport(data common.Data, xlsx *excelize.File, sheet string) int {
	ticketsAvailable := data.Vouchers.AvailableTickets()
	tickets := data.Vouchers.FilterTickets()
	WriteSubHeader(xlsx, sheet, "INGRESSO", strconv.Itoa(offset+1))
	for i, ticket := range ticketsAvailable {
		writeRow(data, xlsx, sheet, tickets.FilterByName(ticket), strconv.Itoa(i+offset+2), ticket, "INGRESSO")
	}
	return offset + len(ticketsAvailable) + 2
}

func genericReport(data common.Data, xlsx *excelize.File, sheet, filter, title string, last int) int {
	ticketsAvailable := data.Vouchers.AvailableFromCategory(filter)
	drinks := data.Vouchers.FilterDrinks().FilterByCategory(filter)
	WriteSubHeader(xlsx, sheet, title, strconv.Itoa(last))
	for i, ticket := range ticketsAvailable {
		writeRow(data, xlsx, sheet, drinks.FilterByName(ticket), strconv.Itoa(1+i+last), ticket, "DRINKS")
	}
	return last + len(ticketsAvailable) + 1
}

func resumo(xlsx *excelize.File, data common.Data, sheet, row, tipo, kind, title string) {
	var drinks types.VouchersList
	list := types.VouchersList(data.Vouchers).FilterByNotStatus("ERROR")
	switch tipo {
	case "CONSUMO":
		drinks = list.FilterDrinks()
	case "TICKET":
		drinks = list.FilterByStatuses("CANCELED", "USED", "AVAILABLE").FilterTickets()
	case "TOTAL":
		drinks = list.
			FilterByStatuses("CANCELED", "USED", "AVAILABLE")
	default:
		drinks = list.FilterDrinks().FilterByCategory(tipo)
	}
	xlsx.SetCellValue(sheet, "A"+row, title)
	xlsx.SetCellValue(sheet, "B"+row, drinks.Size())
	xlsx.SetCellValue(sheet, "D"+row, drinks.FilterByTypes("NORMAL", "TRANSFERED").Size())
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
	if kind == "CONSUMO" {
		liquidFloat = drinks.Accountable().LiquidSumDrinks(data.Club.PercentDrink)
	}
	if kind == "INGRESSO" {
		liquidFloat = drinks.Accountable().LiquidSumTickets(data.Club.PercentTicket, data.Party.AssumeServiceFee)
	}
	if kind == "TOTAL" {
		d := list.FilterTickets().Accountable().LiquidSumTickets(data.Club.PercentTicket, data.Party.AssumeServiceFee)
		t := list.FilterDrinks().Accountable().LiquidSumDrinks(data.Club.PercentDrink)
		liquidFloat = t + d
	}
	xlsx.SetCellValue(sheet, "O"+row, liquidFloat)

	str, _ := common.ReportStyle(xlsx)
	curency, _ := common.CurrencyTable(xlsx)
	xlsx.SetCellStyle(sheet, "A"+row, "B"+row, str)
	xlsx.SetCellStyle(sheet, "D"+row, "H"+row, str)
	xlsx.SetCellStyle(sheet, "J"+row, "L"+row, str)
	xlsx.SetCellStyle(sheet, "N"+row, "O"+row, curency)
}
