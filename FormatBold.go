// Generated using genHelperGo.py
// Thou must not edit this file manually

package color256

import "fmt"

//---------------------------------------------------------------------
// Black
//---------------------------------------------------------------------

// BlackBold is a convenient helper function to print a black and bold text.
// A newline is appended to format by default.
func BlackBold(format string, a ...interface{}) { fmt.Println(StringBlackBold(format,a...)) }

// StringBlackBold is a convenient to get a black and bold string.
func StringBlackBold(format string, a ...interface{}) string { return StringFormat(ColBlack, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Red
//---------------------------------------------------------------------

// RedBold is a convenient helper function to print a red and bold text.
// A newline is appended to format by default.
func RedBold(format string, a ...interface{}) { fmt.Println(StringRedBold(format,a...)) }

// StringRedBold is a convenient to get a red and bold string.
func StringRedBold(format string, a ...interface{}) string { return StringFormat(ColRed, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Green
//---------------------------------------------------------------------

// GreenBold is a convenient helper function to print a green and bold text.
// A newline is appended to format by default.
func GreenBold(format string, a ...interface{}) { fmt.Println(StringGreenBold(format,a...)) }

// StringGreenBold is a convenient to get a green and bold string.
func StringGreenBold(format string, a ...interface{}) string { return StringFormat(ColGreen, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Yellow
//---------------------------------------------------------------------

// YellowBold is a convenient helper function to print a yellow and bold text.
// A newline is appended to format by default.
func YellowBold(format string, a ...interface{}) { fmt.Println(StringYellowBold(format,a...)) }

// StringYellowBold is a convenient to get a yellow and bold string.
func StringYellowBold(format string, a ...interface{}) string { return StringFormat(ColYellow, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Blue
//---------------------------------------------------------------------

// BlueBold is a convenient helper function to print a blue and bold text.
// A newline is appended to format by default.
func BlueBold(format string, a ...interface{}) { fmt.Println(StringBlueBold(format,a...)) }

// StringBlueBold is a convenient to get a blue and bold string.
func StringBlueBold(format string, a ...interface{}) string { return StringFormat(ColBlue, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Magenta
//---------------------------------------------------------------------

// MagentaBold is a convenient helper function to print a magenta and bold text.
// A newline is appended to format by default.
func MagentaBold(format string, a ...interface{}) { fmt.Println(StringMagentaBold(format,a...)) }

// StringMagentaBold is a convenient to get a magenta and bold string.
func StringMagentaBold(format string, a ...interface{}) string { return StringFormat(ColMagenta, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Cyan
//---------------------------------------------------------------------

// CyanBold is a convenient helper function to print a cyan and bold text.
// A newline is appended to format by default.
func CyanBold(format string, a ...interface{}) { fmt.Println(StringCyanBold(format,a...)) }

// StringCyanBold is a convenient to get a cyan and bold string.
func StringCyanBold(format string, a ...interface{}) string { return StringFormat(ColCyan, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// White
//---------------------------------------------------------------------

// WhiteBold is a convenient helper function to print a white and bold text.
// A newline is appended to format by default.
func WhiteBold(format string, a ...interface{}) { fmt.Println(StringWhiteBold(format,a...)) }

// StringWhiteBold is a convenient to get a white and bold string.
func StringWhiteBold(format string, a ...interface{}) string { return StringFormat(ColWhite, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiBlack
//---------------------------------------------------------------------

// HiBlackBold is a convenient helper function to print a hiblack and bold text.
// A newline is appended to format by default.
func HiBlackBold(format string, a ...interface{}) { fmt.Println(StringHiBlackBold(format,a...)) }

// StringHiBlackBold is a convenient to get a hiblack and bold string.
func StringHiBlackBold(format string, a ...interface{}) string { return StringFormat(ColHiBlack, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiRed
//---------------------------------------------------------------------

// HiRedBold is a convenient helper function to print a hired and bold text.
// A newline is appended to format by default.
func HiRedBold(format string, a ...interface{}) { fmt.Println(StringHiRedBold(format,a...)) }

// StringHiRedBold is a convenient to get a hired and bold string.
func StringHiRedBold(format string, a ...interface{}) string { return StringFormat(ColHiRed, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiGreen
//---------------------------------------------------------------------

// HiGreenBold is a convenient helper function to print a higreen and bold text.
// A newline is appended to format by default.
func HiGreenBold(format string, a ...interface{}) { fmt.Println(StringHiGreenBold(format,a...)) }

// StringHiGreenBold is a convenient to get a higreen and bold string.
func StringHiGreenBold(format string, a ...interface{}) string { return StringFormat(ColHiGreen, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiYellow
//---------------------------------------------------------------------

// HiYellowBold is a convenient helper function to print a hiyellow and bold text.
// A newline is appended to format by default.
func HiYellowBold(format string, a ...interface{}) { fmt.Println(StringHiYellowBold(format,a...)) }

// StringHiYellowBold is a convenient to get a hiyellow and bold string.
func StringHiYellowBold(format string, a ...interface{}) string { return StringFormat(ColHiYellow, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiBlue
//---------------------------------------------------------------------

// HiBlueBold is a convenient helper function to print a hiblue and bold text.
// A newline is appended to format by default.
func HiBlueBold(format string, a ...interface{}) { fmt.Println(StringHiBlueBold(format,a...)) }

// StringHiBlueBold is a convenient to get a hiblue and bold string.
func StringHiBlueBold(format string, a ...interface{}) string { return StringFormat(ColHiBlue, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiMagenta
//---------------------------------------------------------------------

// HiMagentaBold is a convenient helper function to print a himagenta and bold text.
// A newline is appended to format by default.
func HiMagentaBold(format string, a ...interface{}) { fmt.Println(StringHiMagentaBold(format,a...)) }

// StringHiMagentaBold is a convenient to get a himagenta and bold string.
func StringHiMagentaBold(format string, a ...interface{}) string { return StringFormat(ColHiMagenta, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiCyan
//---------------------------------------------------------------------

// HiCyanBold is a convenient helper function to print a hicyan and bold text.
// A newline is appended to format by default.
func HiCyanBold(format string, a ...interface{}) { fmt.Println(StringHiCyanBold(format,a...)) }

// StringHiCyanBold is a convenient to get a hicyan and bold string.
func StringHiCyanBold(format string, a ...interface{}) string { return StringFormat(ColHiCyan, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// HiWhite
//---------------------------------------------------------------------

// HiWhiteBold is a convenient helper function to print a hiwhite and bold text.
// A newline is appended to format by default.
func HiWhiteBold(format string, a ...interface{}) { fmt.Println(StringHiWhiteBold(format,a...)) }

// StringHiWhiteBold is a convenient to get a hiwhite and bold string.
func StringHiWhiteBold(format string, a ...interface{}) string { return StringFormat(ColHiWhite, fmt.Sprintf(format,a...),Bold)}

//---------------------------------------------------------------------
// Orange
//---------------------------------------------------------------------

// OrangeBold is a convenient helper function to print a orange and bold text.
// A newline is appended to format by default.
func OrangeBold(format string, a ...interface{}) { fmt.Println(StringOrangeBold(format,a...)) }

// StringOrangeBold is a convenient to get a orange and bold string.
func StringOrangeBold(format string, a ...interface{}) string { return StringFormat(ColOrange, fmt.Sprintf(format,a...),Bold)}

