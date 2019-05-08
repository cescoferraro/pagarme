package common

// Fill dfskjn
type Fill struct {
	Type    string   `json:"type"`
	Color   []string `json:"color"`
	Shading int      `json:"shading"`
	Pattern int      `json:"pattern"`
}

// Pane is a excel panel
type Pane struct {
	Sqref      string `json:"sqref"`
	ActiveCell string `json:"active_cell"`
	PaneName   string `json:"pane"`
}

// PaneStyle is whatever
type PaneStyle struct {
	Freeze      bool   `json:"freeze"`
	Split       bool   `json:"split"`
	XSplit      int    `json:"x_split"`
	YSplit      int    `json:"y_split"`
	TopLeftCell string `json:"top_left_cell"`
	ActivePane  string `json:"active_pane"`
	Panes       []Pane `json:"panes"`
}

//Font fkjsdn
type Font struct {
	Bold   bool   `json:"bold"`
	Italic bool   `json:"italic"`
	Family string `json:"family"`
	Size   int    `json:"size"`
	Color  string `json:"color"`
}

// Boarder vicio
type Boarder struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}

// Alignment sdfjkn
type Alignment struct {
	Vertical    string `json:"vertical"`
	Horizontal  string `json:"horizontal"`
	Center      string `json:"center"`
	ShrinkToFit bool   `json:"shrink_to_fit"`
	WrapText    bool   `json:"wrap_text"`
}

// Style não começou
type Style struct {
	Fill         Fill      `json:"fill"`
	Font         Font      `json:"font"`
	Boarder      []Boarder `json:"border"`
	Alignment    Alignment `json:"alignment"`
	NumberFormat int       `json:"number_format"`
}
