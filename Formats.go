// Generated using genHelperGo.py
// Thou must not edit this file manually

package color256

import "fmt"

//---------------------------------------------------------------------
// Bold
//---------------------------------------------------------------------

// PrintBold is a convenient helper function to print a bold text.
// A newline is appended to format by default.
func PrintBold(format string, a ...interface{}) { fmt.Println(Bold(format,a...)) }

//---------------------------------------------------------------------
// Bright
//---------------------------------------------------------------------

// PrintBright is a convenient helper function to print a bright text.
// A newline is appended to format by default.
func PrintBright(format string, a ...interface{}) { fmt.Println(Bright(format,a...)) }

//---------------------------------------------------------------------
// Faint
//---------------------------------------------------------------------

// PrintFaint is a convenient helper function to print a faint text.
// A newline is appended to format by default.
func PrintFaint(format string, a ...interface{}) { fmt.Println(Faint(format,a...)) }

//---------------------------------------------------------------------
// Italic
//---------------------------------------------------------------------

// PrintItalic is a convenient helper function to print a italic text.
// A newline is appended to format by default.
func PrintItalic(format string, a ...interface{}) { fmt.Println(Italic(format,a...)) }

//---------------------------------------------------------------------
// Underlined
//---------------------------------------------------------------------

// PrintUnderlined is a convenient helper function to print a underlined text.
// A newline is appended to format by default.
func PrintUnderlined(format string, a ...interface{}) { fmt.Println(Underlined(format,a...)) }

//---------------------------------------------------------------------
// Reversed
//---------------------------------------------------------------------

// PrintReversed is a convenient helper function to print a reversed text.
// A newline is appended to format by default.
func PrintReversed(format string, a ...interface{}) { fmt.Println(Reversed(format,a...)) }

