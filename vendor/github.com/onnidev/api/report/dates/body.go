package dates

import (
	"strconv"

	"github.com/onnidev/api/types"
	"github.com/xuri/excelize"
)

const layout = "02/01/2006"

// Body dskjfn
func Body(xlsx *excelize.File, sheet string, all []types.CompleteVoucher) {
	// Total
	list := types.VouchersList(all)
	allTickets := list.FilterTickets().Accountable()
	allDrinks := list.FilterDrinks().Accountable()
	xlsx.SetCellValue(sheet, "A2", "TOTAL")
	xlsx.SetCellValue(sheet, "B2", len(allTickets))
	xlsx.SetCellValue(sheet, "C2", allTickets.Sum())
	xlsx.SetCellValue(sheet, "D2", len(allDrinks))
	xlsx.SetCellValue(sheet, "E2", allDrinks.Sum())
	xlsx.SetCellValue(sheet, "F2", len(allDrinks)+len(allTickets))
	xlsx.SetCellValue(sheet, "G2", append(allDrinks, allTickets...).Sum())
	// Corpo
	hey := list.ConsumedDates()
	for index, day := range hey {
		ii := strconv.Itoa(index + offset + 1)
		dataDrinks := allDrinks.FilterByDate(day)
		dataTickets := allTickets.FilterByDate(day)
		xlsx.SetCellValue(sheet, "A"+ii, day.Format(layout))
		xlsx.SetCellValue(sheet, "B"+ii, len(dataTickets))
		xlsx.SetCellValue(sheet, "C"+ii, dataTickets.Sum())
		xlsx.SetCellValue(sheet, "D"+ii, len(dataDrinks))
		xlsx.SetCellValue(sheet, "E"+ii, dataDrinks.Sum())
		xlsx.SetCellValue(sheet, "F"+ii, len(append(dataTickets, dataDrinks...)))
		xlsx.SetCellValue(sheet, "G"+ii, append(dataTickets, dataDrinks...).Sum())
	}
}
