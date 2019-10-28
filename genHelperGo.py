#!/usr/bin/python3

cols = ["Black",
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
]


with open("FormatNormal.go", "w") as file:
    print("Normal.go")
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

        file.write("// " + col + " is a convenient helper function to print a " + col.lower()+" text.\n")
        file.write("// A newline is appended to format by default.\n")
        file.write("func "+col+"(format string, a ...interface{}) { fmt.Println(String"+col+"(format,a...)) }\n\n")

        file.write("// String" + col +" is a convenient to get a " + col.lower()+" string.\n")
        file.write("func String"+col+"(format string, a ...interface{}) string { return String(Col" +col+", fmt.Sprintf(format,a...))}\n\n")

for f in formats:
    print(f+".go")
    with open("Format"+f+".go","w") as file:
        file.write("// Generated using genHelperGo.py\n")
        file.write("// Thou must not edit this file manually\n\n")
        file.write("package color256\n\n")
        file.write('import "fmt"\n\n')

        for col in cols:
            file.write("//---------------------------------------------------------------------\n")
            file.write("// " + col + "\n")
            file.write("//---------------------------------------------------------------------\n\n")

            file.write("// " + col + f +" is a convenient helper function to print a " + col.lower()+" and " + f.lower()+" text.\n")
            file.write("// A newline is appended to format by default.\n")
            file.write("func " + col + f +"(format string, a ...interface{}) { fmt.Println(String"+col+f+"(format,a...)) }\n\n")

            file.write("// String" + col + f+" is a convenient to get a " + col.lower()+" and " + f.lower()+" string.\n")
            file.write("func String"+col+f+"(format string, a ...interface{}) string { return StringFormat(Col" +col+", fmt.Sprintf(format,a...),"+f+")}\n\n")
