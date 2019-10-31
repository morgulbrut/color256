// Source: https://github.com/esimov/diagram/blob/master/color/color.go

package color256

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mattn/go-isatty"
)

// Colr ansi colors (own type for type checking)
type Colr int

// Formt  text format (own type for type checking)
type Formt int

// Colors
const (
	ColBlack Colr = iota
	ColRed
	ColGreen
	ColYellow
	ColBlue
	ColMagenta
	ColCyan
	ColWhite
	ColHiBlack
	ColHiRed
	ColHiGreen
	ColHiYellow
	ColHiBlue
	ColHiMagenta
	ColHiCyan
	ColHiWhite

	ColOrange   Colr = 208
	ColHiOrange Colr = 214
)

// Formats
const (
	FmtReset Formt = iota
	FmtBold
	FmtFaint
	FmtItalic
	FmtUnderlined
	FmtBlinkSlow
	FmtBlinkRapid
	FmtReversed
	FmtConcealed
	FmtCrossedOut

	FmtFraktur Formt = iota + 19
	FmtDoublyUnderlined

	FmtFramed Formt = iota + 50
	FmtEncircled
	FmtOverlined
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
)

// Aliases
var (
	Bright   = Bold
	Colour   = Color
	BgColour = BgColor
)

// Color returns a color escape string.
func Color(c Colr, str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[39m", c, str)
}

// BgColor returns a backgroundcolor escape string.
func BgColor(c Colr, str string) string {
	return fmt.Sprintf("\x1b[48;5;%dm%s\x1b[49m", c, str)
}

// Format returns a format escape string.
func Format(f Formt, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", f, str)
}

// Bold returns a bold string.
func Bold(format string, a ...interface{}) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[22m", FmtBold, fmt.Sprintf(format, a...))
}

// Faint returns a faint string.
func Faint(format string, a ...interface{}) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[22m", FmtFaint, fmt.Sprintf(format, a...))
}

// Italic returns an italic string.
func Italic(format string, a ...interface{}) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[23m", FmtItalic, fmt.Sprintf(format, a...))
}

// Underlined returns an underlined string.
func Underlined(format string, a ...interface{}) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[24m", FmtUnderlined, fmt.Sprintf(format, a...))
}

// Reversed returns an inversed string.
func Reversed(format string, a ...interface{}) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[27m", FmtReversed, fmt.Sprintf(format, a...))
}

// PrintRandom  is a convenient helper function to print with random foreground.
// A newline is appended to format by default.
func PrintRandom(format string, a ...interface{}) { fmt.Println(Random(format, a...)) }

// Random  is a convenient helper function to print with black foreground.
// A newline is appended to format by default.
func Random(format string, a ...interface{}) string {
	return Color(Colr(RandomColor(180, 231)), fmt.Sprintf(format, a...))
}

// RandomColor returns a random  color number.
func RandomColor(min, max int) int {
	return rand.Intn(max-min) + min
}

// Init random number generator (Quick and dirty Windows hack)
func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// List all the colors with its color code
func List() {
	for i := 0; i < 256; i++ {
		fmt.Println(Color(Colr(i), fmt.Sprintf("%d", i)))
	}
}
