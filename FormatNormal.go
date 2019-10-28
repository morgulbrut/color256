// Generated using genHelperGo.py
// Thou must not edit this file manually

package color256

import "fmt"

//---------------------------------------------------------------------
// Black
//---------------------------------------------------------------------

// Black is a convenient helper function to print a black text.
// A newline is appended to format by default.
func Black(format string, a ...interface{}) { fmt.Println(StringBlack(format, a...)) }

// StringBlack is a convenient to get a black string.
func StringBlack(format string, a ...interface{}) string {
	return String(ColBlack, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Red
//---------------------------------------------------------------------

// Red is a convenient helper function to print a red text.
// A newline is appended to format by default.
func Red(format string, a ...interface{}) { fmt.Println(StringRed(format, a...)) }

// StringRed is a convenient to get a red string.
func StringRed(format string, a ...interface{}) string {
	return String(ColRed, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Green
//---------------------------------------------------------------------

// Green is a convenient helper function to print a green text.
// A newline is appended to format by default.
func Green(format string, a ...interface{}) { fmt.Println(StringGreen(format, a...)) }

// StringGreen is a convenient to get a green string.
func StringGreen(format string, a ...interface{}) string {
	return String(ColGreen, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Yellow
//---------------------------------------------------------------------

// Yellow is a convenient helper function to print a yellow text.
// A newline is appended to format by default.
func Yellow(format string, a ...interface{}) { fmt.Println(StringYellow(format, a...)) }

// StringYellow is a convenient to get a yellow string.
func StringYellow(format string, a ...interface{}) string {
	return String(ColYellow, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Blue
//---------------------------------------------------------------------

// Blue is a convenient helper function to print a blue text.
// A newline is appended to format by default.
func Blue(format string, a ...interface{}) { fmt.Println(StringBlue(format, a...)) }

// StringBlue is a convenient to get a blue string.
func StringBlue(format string, a ...interface{}) string {
	return String(ColBlue, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Magenta
//---------------------------------------------------------------------

// Magenta is a convenient helper function to print a magenta text.
// A newline is appended to format by default.
func Magenta(format string, a ...interface{}) { fmt.Println(StringMagenta(format, a...)) }

// StringMagenta is a convenient to get a magenta string.
func StringMagenta(format string, a ...interface{}) string {
	return String(ColMagenta, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Cyan
//---------------------------------------------------------------------

// Cyan is a convenient helper function to print a cyan text.
// A newline is appended to format by default.
func Cyan(format string, a ...interface{}) { fmt.Println(StringCyan(format, a...)) }

// StringCyan is a convenient to get a cyan string.
func StringCyan(format string, a ...interface{}) string {
	return String(ColCyan, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// White
//---------------------------------------------------------------------

// White is a convenient helper function to print a white text.
// A newline is appended to format by default.
func White(format string, a ...interface{}) { fmt.Println(StringWhite(format, a...)) }

// StringWhite is a convenient to get a white string.
func StringWhite(format string, a ...interface{}) string {
	return String(ColWhite, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiBlack
//---------------------------------------------------------------------

// HiBlack is a convenient helper function to print a hiblack text.
// A newline is appended to format by default.
func HiBlack(format string, a ...interface{}) { fmt.Println(StringHiBlack(format, a...)) }

// StringHiBlack is a convenient to get a hiblack string.
func StringHiBlack(format string, a ...interface{}) string {
	return String(ColHiBlack, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiRed
//---------------------------------------------------------------------

// HiRed is a convenient helper function to print a hired text.
// A newline is appended to format by default.
func HiRed(format string, a ...interface{}) { fmt.Println(StringHiRed(format, a...)) }

// StringHiRed is a convenient to get a hired string.
func StringHiRed(format string, a ...interface{}) string {
	return String(ColHiRed, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiGreen
//---------------------------------------------------------------------

// HiGreen is a convenient helper function to print a higreen text.
// A newline is appended to format by default.
func HiGreen(format string, a ...interface{}) { fmt.Println(StringHiGreen(format, a...)) }

// StringHiGreen is a convenient to get a higreen string.
func StringHiGreen(format string, a ...interface{}) string {
	return String(ColHiGreen, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiYellow
//---------------------------------------------------------------------

// HiYellow is a convenient helper function to print a hiyellow text.
// A newline is appended to format by default.
func HiYellow(format string, a ...interface{}) { fmt.Println(StringHiYellow(format, a...)) }

// StringHiYellow is a convenient to get a hiyellow string.
func StringHiYellow(format string, a ...interface{}) string {
	return String(ColHiYellow, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiBlue
//---------------------------------------------------------------------

// HiBlue is a convenient helper function to print a hiblue text.
// A newline is appended to format by default.
func HiBlue(format string, a ...interface{}) { fmt.Println(StringHiBlue(format, a...)) }

// StringHiBlue is a convenient to get a hiblue string.
func StringHiBlue(format string, a ...interface{}) string {
	return String(ColHiBlue, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiMagenta
//---------------------------------------------------------------------

// HiMagenta is a convenient helper function to print a himagenta text.
// A newline is appended to format by default.
func HiMagenta(format string, a ...interface{}) { fmt.Println(StringHiMagenta(format, a...)) }

// StringHiMagenta is a convenient to get a himagenta string.
func StringHiMagenta(format string, a ...interface{}) string {
	return String(ColHiMagenta, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiCyan
//---------------------------------------------------------------------

// HiCyan is a convenient helper function to print a hicyan text.
// A newline is appended to format by default.
func HiCyan(format string, a ...interface{}) { fmt.Println(StringHiCyan(format, a...)) }

// StringHiCyan is a convenient to get a hicyan string.
func StringHiCyan(format string, a ...interface{}) string {
	return String(ColHiCyan, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// HiWhite
//---------------------------------------------------------------------

// HiWhite is a convenient helper function to print a hiwhite text.
// A newline is appended to format by default.
func HiWhite(format string, a ...interface{}) { fmt.Println(StringHiWhite(format, a...)) }

// StringHiWhite is a convenient to get a hiwhite string.
func StringHiWhite(format string, a ...interface{}) string {
	return String(ColHiWhite, fmt.Sprintf(format, a...))
}

//---------------------------------------------------------------------
// Orange
//---------------------------------------------------------------------

// Orange is a convenient helper function to print a orange text.
// A newline is appended to format by default.
func Orange(format string, a ...interface{}) { fmt.Println(StringOrange(format, a...)) }

// StringOrange is a convenient to get a orange string.
func StringOrange(format string, a ...interface{}) string {
	return String(ColOrange, fmt.Sprintf(format, a...))
}
