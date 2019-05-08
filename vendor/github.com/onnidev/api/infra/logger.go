package infra

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
	jww "github.com/spf13/jwalterweatherman"
)

// Logger TODO: NEEDS COMMENT INFO
type Logger struct {
	Title string
	Color color.Attribute
	Time  time.Time
}

var (
	acceptableColors []color.Attribute
)

// Now TODO: NEEDS COMMENT INFO
func (block Logger) Now() Logger {
	block.Time = time.Now()
	return block
}
func init() {
	renew()
	jww.SetStdoutThreshold(jww.LevelTrace)
}

// NewLogger TODO: NEEDS COMMENT INFO
func NewLogger(title string) Logger {
	return Logger{Title: title, Color: getColor()}
}

// NewColor TODO: NEEDS COMMENT INFO
func NewColor(title string, color color.Attribute) Logger {
	for u, cor := range acceptableColors {
		if color == cor {
			removeColor(u)
		}
	}
	return Logger{Title: title, Color: color}
}

func removeColor(random int) {
	acceptableColors = append(acceptableColors[:random], acceptableColors[random+1:]...)
	if len(acceptableColors) == 0 {
		renew()
	}
}

func getColor() color.Attribute {
	random := rand.Intn(len(acceptableColors))
	defer removeColor(random)
	return acceptableColors[random]
}

// Print TODO: NEEDS COMMENT INFO
func (block Logger) Print(message ...interface{}) {
	msg := block.Header()
	jww.TRACE.Printf(msg+"%v\n", message)
}

// Printf TODO: NEEDS COMMENT INFO
func (block Logger) Printf(format string, v ...interface{}) {
	msg := block.Header()
	jww.TRACE.Printf(msg+format, v)
}

// Header TODO: NEEDS COMMENT INFO
func (block Logger) Header() string {
	colore := color.New(block.Color).SprintFunc()
	return "[" + colore(block.Title) + "] "
}

// WIthTimer TODO: NEEDS COMMENT INFO
func (block Logger) WIthTimer() string {
	colore := color.New(block.Color).SprintFunc()
	return colore("+") + colore(time.Since(block.Time))
}

func renew() {
	acceptableColors = []color.
		Attribute{
		//color.FgBlack,
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		//color.FgWhite,
		color.BgBlack,
		color.BgRed,
		color.BgGreen,
		color.BgYellow,
		color.BgBlue,
		color.BgMagenta,
		color.BgCyan,
		//color.BgWhite,
		color.FgHiBlack,
		color.FgHiRed,
		color.FgHiGreen,
		color.FgHiYellow,
		color.FgHiBlue,
		color.FgHiMagenta,
		color.FgHiCyan,
		//color.FgHiWhite,
	}
}
