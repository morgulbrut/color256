// Generated using genHelperGo.py
// Thou must not edit this file manually

package color256

import "fmt"

// Bold returns a bold string.
func Bold(format string, a ...interface{}) string {
	return fmt.Sprintf("[%dm%s[22m", FmtBold, fmt.Sprintf(format, a...))
}

// Faint returns a faint string.
func Faint(format string, a ...interface{}) string {
	return fmt.Sprintf("[%dm%s[22m", FmtFaint, fmt.Sprintf(format, a...))
}

// Italic returns an italic string.
func Italic(format string, a ...interface{}) string {
	return fmt.Sprintf("[%dm%s[23m", FmtItalic, fmt.Sprintf(format, a...))
}

// Underlined returns an underlined string.
func Underlined(format string, a ...interface{}) string {
	return fmt.Sprintf("[%dm%s[24m", FmtUnderlined, fmt.Sprintf(format, a...))
}

// Reversed returns an inversed string.
func Reversed(format string, a ...interface{}) string {
	return fmt.Sprintf("[%dm%s[27m", FmtReversed, fmt.Sprintf(format, a...))
}
// PrintBold is a convenient helper function to print a bold text.
// A newline is appended to format by default.
func PrintBold(format string, a ...interface{}) { fmt.Println(Bold(format,a...)) }

// PrintBright is a convenient helper function to print a bright text.
// A newline is appended to format by default.
func PrintBright(format string, a ...interface{}) { fmt.Println(Bright(format,a...)) }

// PrintFaint is a convenient helper function to print a faint text.
// A newline is appended to format by default.
func PrintFaint(format string, a ...interface{}) { fmt.Println(Faint(format,a...)) }

// PrintItalic is a convenient helper function to print a italic text.
// A newline is appended to format by default.
func PrintItalic(format string, a ...interface{}) { fmt.Println(Italic(format,a...)) }

// PrintUnderlined is a convenient helper function to print a underlined text.
// A newline is appended to format by default.
func PrintUnderlined(format string, a ...interface{}) { fmt.Println(Underlined(format,a...)) }

// PrintReversed is a convenient helper function to print a reversed text.
// A newline is appended to format by default.
func PrintReversed(format string, a ...interface{}) { fmt.Println(Reversed(format,a...)) }

