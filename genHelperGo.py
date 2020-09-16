#!/usr/bin/python3

cols = [
    "Black",
    "Red",
    "Green",
    "Yellow",
    "Blue",
    "Magenta",
    "Cyan",
    "White",
    "HiBlack",
    "HiRed",
    "HiGreen",
    "HiYellow",
    "HiBlue",
    "HiMagenta",
    "HiCyan",
    "HiWhite",
    "Orange",
    "HiOrange",
    "Pink",
    "HiPink",
]

formats = [
    "Bold",
    "Bright",
    "Faint",
    "Italic",
    "Underlined",
    "Reversed",
]


with open("Colors.go", "w") as file:
    print("Colors.go")
    file.write("// Generated using genHelperGo.py\n")
    file.write("// Thou must not edit this file manually\n\n")
    file.write("package color256\n\n")
    file.write('import "fmt"\n\n')

    for col in cols:
        file.write(
            "//---------------------------------------------------------------------\n")
        file.write("// " + col + "\n")
        file.write(
            "//---------------------------------------------------------------------\n\n")

        file.write("// Print" + col + " is a convenient helper function to print a " + col.lower()+" text.\n")
        file.write("// A newline is appended to format by default.\n")
        file.write("func Print"+col+"(format string, a ...interface{}) { fmt.Println("+col+"(format,a...)) }\n\n")

        file.write("// " + col +" is a convenient to get a " + col.lower()+" string.\n")
        file.write("func "+col+"(format string, a ...interface{}) string { return Color(Col" +col+", fmt.Sprintf(format,a...))}\n\n")

with open("BgColors.go", "w") as file:
    print("BgColors.go")
    file.write("// Generated using genHelperGo.py\n")
    file.write("// Thou must not edit this file manually\n\n")
    file.write("package color256\n\n")
    file.write('import "fmt"\n\n')

    for col in cols:
        file.write(
            "//---------------------------------------------------------------------\n")
        file.write("// " + col + "\n")
        file.write(
            "//---------------------------------------------------------------------\n\n")

        file.write("// PrintBg" + col + " is a convenient helper function to print a " + col.lower()+" text.\n")
        file.write("// A newline is appended to format by default.\n")
        file.write("func PrintBg"+col+"(format string, a ...interface{}) { fmt.Println("+col+"(format,a...)) }\n\n")

        file.write("// Bg" + col +" is a convenient to get a " + col.lower()+" string.\n")
        file.write("func Bg"+col+"(format string, a ...interface{}) string { return BgColor(Col" +col+", fmt.Sprintf(format,a...))}\n\n")


with open("Formats.go", "w") as file:
    print("Formats.go")
    
    file.write("// Generated using genHelperGo.py\n")
    file.write("// Thou must not edit this file manually\n\n")
    file.write("package color256\n\n")
    file.write('import "fmt"\n\n')

    funcs = '''// Bold returns a bold string.
func Bold(format string, a ...interface{}) string {
	return fmt.Sprintf("\\x1b[%dm%s\\x1b[22m", FmtBold, fmt.Sprintf(format, a...))
}

// Faint returns a faint string.
func Faint(format string, a ...interface{}) string {
	return fmt.Sprintf("\\x1b[%dm%s\\x1b[22m", FmtFaint, fmt.Sprintf(format, a...))
}

// Italic returns an italic string.
func Italic(format string, a ...interface{}) string {
	return fmt.Sprintf("\\x1b[%dm%s\\x1b[23m", FmtItalic, fmt.Sprintf(format, a...))
}

// Underlined returns an underlined string.
func Underlined(format string, a ...interface{}) string {
	return fmt.Sprintf("\\x1b[%dm%s\\x1b[24m", FmtUnderlined, fmt.Sprintf(format, a...))
}

// Reversed returns an inversed string.
func Reversed(format string, a ...interface{}) string {
	return fmt.Sprintf("\\x1b[%dm%s\\x1b[27m", FmtReversed, fmt.Sprintf(format, a...))
}
'''
    file.write(funcs)

    for f in formats:   
        file.write("// Print" + f +" is a convenient helper function to print a " + f.lower()+" text.\n")
        file.write("// A newline is appended to format by default.\n")
        file.write("func Print" + f +"(format string, a ...interface{}) { fmt.Println("+f+"(format,a...)) }\n\n")


with open("color256_test.go", "w") as file:
    print("color256_test.go")
    test = '''
package color256

import (
	"fmt"
	"testing"
    "os"
)
'''   
    file.write("// Generated using genHelperGo.py\n")
    file.write("// Thou must not edit this file manually\n\n")
    file.write(test)
    
    for col in cols:
        file.write("func Test"+col+"(t *testing.T) {\n")
        for f in formats:
            file.write("fmt.Println("+col+"(\""+col+" and \"+"+f +"(\""+f+"\")  + \" and "+col+"\"))\n")
        file.write("}\n\n")
    
    file.write("func TestConvFuncs (t *testing.T) {\n")
    for col in cols:
        file.write("Print"+col+"(\""+col+"\")\n")
        file.write("PrintBg"+col+"(\""+col+"\")\n")
    file.write("}\n\n")

    file.write('''// Prints out if the enviroment variable NO_COLOR is set
func TestNO_COLOR(t *testing.T){
    _, noColor := os.LookupEnv("NO_COLOR")
	if noColor {
		fmt.Println("No color")
	} else {
		fmt.Println("color")
	}
}''')
