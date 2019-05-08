package users

import (
	"sort"
	"strconv"
	"unicode"

	"github.com/onnidev/api/report/common"
	"github.com/xuri/excelize"
)

// Report is sickest
func Report(xlsx *excelize.File, data common.Data) {
	sheet := "Sheet4"
	xlsx.NewSheet(4, sheet)
	// sort.Sort(data.Vouchers)
	xlsx.SetCellValue(sheet, "A1", "EMAIL")
	xlsx.SetCellValue(sheet, "B1", "TICKET")
	xlsx.SetCellValue(sheet, "C1", "TICKET R$")
	xlsx.SetCellValue(sheet, "D1", "DRINKS")
	xlsx.SetCellValue(sheet, "E1", "DRINKS R$")
	xlsx.SetCellValue(sheet, "F1", "TOTAL")
	xlsx.SetCellValue(sheet, "G1", "TOTAL R$")
	xlsx.SetColWidth(sheet, "A", "A", 30)
	xlsx.SetColWidth(sheet, "B", "G", 15)
	gap := 2
	xlsx.SetCellValue(sheet, "A2", "Transações por usuário/Ticket médio geral")
	xlsx.SetCellValue(sheet, "B2", data.Vouchers.FilterTickets().Accountable().Size())
	xlsx.SetCellValue(sheet, "C2", data.Vouchers.FilterTickets().Accountable().Sum())
	xlsx.SetCellValue(sheet, "D2", data.Vouchers.FilterDrinks().Accountable().Size())
	xlsx.SetCellValue(sheet, "E2", data.Vouchers.FilterDrinks().Accountable().Sum())
	xlsx.SetCellValue(sheet, "F2", data.Vouchers.Accountable().Size())
	xlsx.SetCellValue(sheet, "G2", data.Vouchers.Accountable().Sum())

	customers := ByCase(data.Vouchers.AvailableCustomers())
	sort.Sort(customers)

	regular, _ := common.Table(xlsx)
	black, _ := common.BlackTable(xlsx)
	currency, _ := common.CurrencyTable(xlsx)
	items, _ := common.HeaderStyle(xlsx)
	xlsx.SetCellStyle(sheet, "A1", "G1", black)
	xlsx.SetCellStyle(sheet, "A2", "A"+strconv.Itoa(gap), items)
	xlsx.SetCellStyle(sheet, "B"+strconv.Itoa(2), "B"+strconv.Itoa(len(customers)+10), regular)
	xlsx.SetCellStyle(sheet, "C"+strconv.Itoa(2), "C"+strconv.Itoa(len(customers)+10), currency)
	xlsx.SetCellStyle(sheet, "D"+strconv.Itoa(2), "D"+strconv.Itoa(len(customers)+10), regular)
	xlsx.SetCellStyle(sheet, "E"+strconv.Itoa(2), "E"+strconv.Itoa(len(customers)+10), currency)
	xlsx.SetCellStyle(sheet, "F"+strconv.Itoa(2), "F"+strconv.Itoa(len(customers)+10), regular)
	xlsx.SetCellStyle(sheet, "G"+strconv.Itoa(2), "G"+strconv.Itoa(len(customers)+10), currency)
	xlsx.SetCellStyle(sheet, "A"+strconv.Itoa(gap+1), "G"+strconv.Itoa(len(customers)+10), regular)
	for index, mail := range customers {
		customer := data.Vouchers.
			Accountable().
			FilterByCustomerMail(mail)
		xlsx.SetCellValue(sheet, "A"+strconv.Itoa(1+index+gap), mail)
		xlsx.SetCellValue(sheet, "B"+strconv.Itoa(1+index+gap), customer.FilterTickets().Size())
		xlsx.SetCellValue(sheet, "C"+strconv.Itoa(1+index+gap), customer.FilterTickets().Sum())
		xlsx.SetCellValue(sheet, "D"+strconv.Itoa(1+index+gap), customer.FilterDrinks().Size())
		xlsx.SetCellValue(sheet, "E"+strconv.Itoa(1+index+gap), customer.FilterDrinks().Sum())
		xlsx.SetCellValue(sheet, "F"+strconv.Itoa(1+index+gap), customer.Size())
		xlsx.SetCellValue(sheet, "G"+strconv.Itoa(1+index+gap), customer.Sum())
	}

	for i := 0; i < gap+len(customers); i++ {
		xlsx.SetRowHeight(sheet, i, 25)
	}
	xlsx.SetSheetName(sheet, "Análise de Ticket")
}

// ByCase dfkjgndfg
type ByCase []string

func (s ByCase) Len() int      { return len(s) }
func (s ByCase) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByCase) Less(i, j int) bool {
	iRunes := []rune(s[i])
	jRunes := []rune(s[j])

	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		// the lowercase runes are the same, so compare the original
		if ir != jr {
			return ir < jr
		}
	}

	return false
}
