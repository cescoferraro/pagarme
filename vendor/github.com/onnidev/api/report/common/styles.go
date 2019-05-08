package common

import (
	"encoding/json"
	"strconv"

	"github.com/xuri/excelize"
)

var whiteFont = Font{
	Bold:   false,
	Italic: true,
	Family: "Arial",
	Size:   12,
	Color:  "#Ffffff"}

var defaultFont = Font{
	Bold:   false,
	Italic: true,
	Family: "Arial",
	Size:   12,
	Color:  "#000000"}

var center = Alignment{
	Vertical:    "center",
	Horizontal:  "center",
	ShrinkToFit: true,
	WrapText:    false}

var subHeaderBackground = Fill{
	Type:    "pattern",
	Color:   []string{"#D6DCE4"},
	Pattern: 1}

var headerBackground = Fill{
	Type:    "pattern",
	Color:   []string{"#222B35"},
	Pattern: 1}

var border = []Boarder{
	Boarder{Type: "left", Color: "#000000", Style: 1},
	Boarder{Type: "right", Color: "#000000", Style: 1},
	Boarder{Type: "top", Color: "#000000", Style: 1},
	Boarder{Type: "bottom", Color: "#000000", Style: 1}}

var allBlack = Fill{
	Type:    "gradient",
	Color:   []string{"#000000", "#000000"},
	Shading: 1}

var regularBackground = Fill{
	Type:    "gradient",
	Color:   []string{"#FFFFFF", "#F5f5f5"},
	Shading: 1}

// SubHeaderStyle is a type
func SubHeaderStyle(xlsx *excelize.File) (int, error) {
	var headerStyle = Style{
		Fill:         subHeaderBackground,
		Font:         defaultFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 1}
	str, _ := json.Marshal(headerStyle)
	return xlsx.NewStyle(string(str))
}

// HeaderStyle is a type
func HeaderStyle(xlsx *excelize.File) (int, error) {
	var headerStyle = Style{
		Fill:         headerBackground,
		Font:         whiteFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 1}
	str, _ := json.Marshal(headerStyle)
	return xlsx.NewStyle(string(str))
}

// ReportStyle is a type
func ReportStyle(xlsx *excelize.File) (int, error) {
	var reportStyle = Style{
		Fill:         regularBackground,
		Font:         defaultFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 1}
	str, _ := json.Marshal(reportStyle)
	return xlsx.NewStyle(string(str))
}

// Table is a type
func Table(xlsx *excelize.File) (int, error) {
	var tableHeaderStyle = Style{
		Fill:         regularBackground,
		Font:         defaultFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 1}
	str, _ := json.Marshal(tableHeaderStyle)
	return xlsx.NewStyle(string(str))
}

// CurrencyTable is a type
func CurrencyTable(xlsx *excelize.File) (int, error) {
	var currencyStyle = Style{
		Fill:         regularBackground,
		Font:         defaultFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 353,
	}
	str, _ := json.Marshal(currencyStyle)
	return xlsx.NewStyle(string(str))
}

// BlackTable is a type
func BlackTable(xlsx *excelize.File) (int, error) {
	var dateStyle = Style{
		Fill:         allBlack,
		Font:         whiteFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 1}
	str, _ := json.Marshal(dateStyle)
	return xlsx.NewStyle(string(str))
}

// DateTable is a type
func DateTable(xlsx *excelize.File) (int, error) {

	var dateStyle = Style{
		Fill:         regularBackground,
		Font:         defaultFont,
		Boarder:      border,
		Alignment:    center,
		NumberFormat: 22,
	}

	str, _ := json.Marshal(dateStyle)
	return xlsx.NewStyle(string(str))
}

// SetPanes is a type
func SetPanes(xlsx *excelize.File) {
	ee, _ := json.Marshal(change)
	xlsx.SetPanes("Sheet1", string(ee))
}

// VouchersSetPanes is a type
func VouchersSetPanes(xlsx *excelize.File) {
	ee, _ := json.Marshal(vouchersPane)
	xlsx.SetPanes("Sheet2", string(ee))
}

var change = func(offset int) PaneStyle {
	return PaneStyle{
		Freeze:      true,
		Split:       false,
		XSplit:      1,
		YSplit:      offset,
		TopLeftCell: "B" + strconv.Itoa(offset+1),
		ActivePane:  "bottomLeft",
		Panes: []Pane{Pane{
			Sqref:      "A!!:XFD!!",
			ActiveCell: "A11",
			PaneName:   "bottomLeft"}}}
}

var vouchersPane = func(offset int) PaneStyle {
	return PaneStyle{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      2,
		TopLeftCell: "A3",
		ActivePane:  "vouchers",
		Panes: []Pane{{
			Sqref:      "A!!:XFD!!",
			ActiveCell: "A11",
			PaneName:   "vouchers"}}}
}
