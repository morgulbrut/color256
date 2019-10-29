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
]

formats = [
    "Bold",
    "Faint",
    "Italic",
    "Underlined",
    "BlinkSlow",
    "BlinkRapid",
    "Reversed",
    "Concealed",
    "CrossedOut",

    "Fraktur",
    "DoublyUnderlined",

    "Framed",
    "Encircled",
    "Overlined",

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

with open("Formats.go", "w") as file:
    print("Formats.go")
    
    file.write("// Generated using genHelperGo.py\n")
    file.write("// Thou must not edit this file manually\n\n")
    file.write("package color256\n\n")
    file.write('import "fmt"\n\n')

    for f in formats:   
        file.write("//---------------------------------------------------------------------\n")
        file.write("// " + f + "\n")
        file.write("//---------------------------------------------------------------------\n\n")

        file.write("// Print" + f +" is a convenient helper function to print a " + f.lower()+" text.\n")
        file.write("// A newline is appended to format by default.\n")
        file.write("func Print" + f +"(format string, a ...interface{}) { fmt.Println("+f+"(format,a...)) }\n\n")

        file.write("//" +f+ " is a convenient to get a " + f.lower()+" string.\n")
        file.write("func " +f+ "(format string, a ...interface{}) string { return Format(Fmt" +f+", fmt.Sprintf(format,a...))}\n\n")
