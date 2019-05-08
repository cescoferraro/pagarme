package dates

import (
	"encoding/json"
	"strconv"

	"github.com/onnidev/api/types"
	"github.com/xuri/excelize"
)

// Charts sdkjfn
func Charts(xlsx *excelize.File, sheet string, all []types.CompleteVoucher) {
	list := types.VouchersList(all)
	hey := list.ConsumedDates()
	last := strconv.Itoa(len(hey) + 2)
	infame := strconv.Itoa(offset + len(hey) + 2)
	gg, _ := json.Marshal(hrapg(last))
	xlsx.AddChart(sheet, "A"+infame, string(gg))
	hh, _ := json.Marshal(amount(last))
	xlsx.AddChart(sheet, "E"+infame, string(hh))
}

func amount(last string) Chart {
	return Chart{
		Title: "Fruit 3D Line CHart",
		Type:  "bar3D",
		Series: []Serie{
			{Name: "=Datas!$C$1", Categories: "=Datas!$A$3:$A$" + last, Values: "=Datas!$C$3:$C$" + last},
			{Name: "=Datas!$E$1", Categories: "=Datas!$A$3:$A$" + last, Values: "=Datas!$E$3:$E$" + last},
			// {Name: "=Datas!$G$1", Categories: "=Datas!$A$2:$A$9", Values: "=Datas!$G$2:$G$9"},
		},
	}
}

func hrapg(last string) Chart {
	return Chart{
		Title: "Fruit 3D Line CHart",
		Type:  "bar3D",
		Series: []Serie{
			{Name: "=Datas!$B$1", Categories: "=Datas!$A$3:$A$" + last, Values: "=Datas!$B$3:$B$" + last},
			{Name: "=Datas!$D$1", Categories: "=Datas!$A$3:$A$" + last, Values: "=Datas!$D$3:$D$" + last},
			// {Name: "=Datas!$F$1", Categories: "=Datas!$A$2:$A$9", Values: "=Datas!$F$2:$F$9"},
		},
	}
}

// Chart skdjfgn
type Chart struct {
	Type   string  `json:"type"`
	Series []Serie `json:"series"`
	Title  string  `json:"title"`
}

// Serie sdjkfn
type Serie struct {
	Name       string `json:"name"`
	Categories string `json:"categories"`
	Values     string `json:"values"`
}
