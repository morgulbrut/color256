// Generated using genHelperGo.py
// Thou must not edit this file manually

package color256

import (
	"fmt"
	"os"
	"testing"
)

func TestBlack(t *testing.T) {
	fmt.Println(Black("Black and " + Bold("Bold") + " and Black"))
	fmt.Println(Black("Black and " + Bright("Bright") + " and Black"))
	fmt.Println(Black("Black and " + Faint("Faint") + " and Black"))
	fmt.Println(Black("Black and " + Italic("Italic") + " and Black"))
	fmt.Println(Black("Black and " + Underlined("Underlined") + " and Black"))
	fmt.Println(Black("Black and " + Reversed("Reversed") + " and Black"))
}

func TestRed(t *testing.T) {
	fmt.Println(Red("Red and " + Bold("Bold") + " and Red"))
	fmt.Println(Red("Red and " + Bright("Bright") + " and Red"))
	fmt.Println(Red("Red and " + Faint("Faint") + " and Red"))
	fmt.Println(Red("Red and " + Italic("Italic") + " and Red"))
	fmt.Println(Red("Red and " + Underlined("Underlined") + " and Red"))
	fmt.Println(Red("Red and " + Reversed("Reversed") + " and Red"))
}

func TestGreen(t *testing.T) {
	fmt.Println(Green("Green and " + Bold("Bold") + " and Green"))
	fmt.Println(Green("Green and " + Bright("Bright") + " and Green"))
	fmt.Println(Green("Green and " + Faint("Faint") + " and Green"))
	fmt.Println(Green("Green and " + Italic("Italic") + " and Green"))
	fmt.Println(Green("Green and " + Underlined("Underlined") + " and Green"))
	fmt.Println(Green("Green and " + Reversed("Reversed") + " and Green"))
}

func TestYellow(t *testing.T) {
	fmt.Println(Yellow("Yellow and " + Bold("Bold") + " and Yellow"))
	fmt.Println(Yellow("Yellow and " + Bright("Bright") + " and Yellow"))
	fmt.Println(Yellow("Yellow and " + Faint("Faint") + " and Yellow"))
	fmt.Println(Yellow("Yellow and " + Italic("Italic") + " and Yellow"))
	fmt.Println(Yellow("Yellow and " + Underlined("Underlined") + " and Yellow"))
	fmt.Println(Yellow("Yellow and " + Reversed("Reversed") + " and Yellow"))
}

func TestBlue(t *testing.T) {
	fmt.Println(Blue("Blue and " + Bold("Bold") + " and Blue"))
	fmt.Println(Blue("Blue and " + Bright("Bright") + " and Blue"))
	fmt.Println(Blue("Blue and " + Faint("Faint") + " and Blue"))
	fmt.Println(Blue("Blue and " + Italic("Italic") + " and Blue"))
	fmt.Println(Blue("Blue and " + Underlined("Underlined") + " and Blue"))
	fmt.Println(Blue("Blue and " + Reversed("Reversed") + " and Blue"))
}

func TestMagenta(t *testing.T) {
	fmt.Println(Magenta("Magenta and " + Bold("Bold") + " and Magenta"))
	fmt.Println(Magenta("Magenta and " + Bright("Bright") + " and Magenta"))
	fmt.Println(Magenta("Magenta and " + Faint("Faint") + " and Magenta"))
	fmt.Println(Magenta("Magenta and " + Italic("Italic") + " and Magenta"))
	fmt.Println(Magenta("Magenta and " + Underlined("Underlined") + " and Magenta"))
	fmt.Println(Magenta("Magenta and " + Reversed("Reversed") + " and Magenta"))
}

func TestCyan(t *testing.T) {
	fmt.Println(Cyan("Cyan and " + Bold("Bold") + " and Cyan"))
	fmt.Println(Cyan("Cyan and " + Bright("Bright") + " and Cyan"))
	fmt.Println(Cyan("Cyan and " + Faint("Faint") + " and Cyan"))
	fmt.Println(Cyan("Cyan and " + Italic("Italic") + " and Cyan"))
	fmt.Println(Cyan("Cyan and " + Underlined("Underlined") + " and Cyan"))
	fmt.Println(Cyan("Cyan and " + Reversed("Reversed") + " and Cyan"))
}

func TestWhite(t *testing.T) {
	fmt.Println(White("White and " + Bold("Bold") + " and White"))
	fmt.Println(White("White and " + Bright("Bright") + " and White"))
	fmt.Println(White("White and " + Faint("Faint") + " and White"))
	fmt.Println(White("White and " + Italic("Italic") + " and White"))
	fmt.Println(White("White and " + Underlined("Underlined") + " and White"))
	fmt.Println(White("White and " + Reversed("Reversed") + " and White"))
}

func TestHiBlack(t *testing.T) {
	fmt.Println(HiBlack("HiBlack and " + Bold("Bold") + " and HiBlack"))
	fmt.Println(HiBlack("HiBlack and " + Bright("Bright") + " and HiBlack"))
	fmt.Println(HiBlack("HiBlack and " + Faint("Faint") + " and HiBlack"))
	fmt.Println(HiBlack("HiBlack and " + Italic("Italic") + " and HiBlack"))
	fmt.Println(HiBlack("HiBlack and " + Underlined("Underlined") + " and HiBlack"))
	fmt.Println(HiBlack("HiBlack and " + Reversed("Reversed") + " and HiBlack"))
}

func TestHiRed(t *testing.T) {
	fmt.Println(HiRed("HiRed and " + Bold("Bold") + " and HiRed"))
	fmt.Println(HiRed("HiRed and " + Bright("Bright") + " and HiRed"))
	fmt.Println(HiRed("HiRed and " + Faint("Faint") + " and HiRed"))
	fmt.Println(HiRed("HiRed and " + Italic("Italic") + " and HiRed"))
	fmt.Println(HiRed("HiRed and " + Underlined("Underlined") + " and HiRed"))
	fmt.Println(HiRed("HiRed and " + Reversed("Reversed") + " and HiRed"))
}

func TestHiGreen(t *testing.T) {
	fmt.Println(HiGreen("HiGreen and " + Bold("Bold") + " and HiGreen"))
	fmt.Println(HiGreen("HiGreen and " + Bright("Bright") + " and HiGreen"))
	fmt.Println(HiGreen("HiGreen and " + Faint("Faint") + " and HiGreen"))
	fmt.Println(HiGreen("HiGreen and " + Italic("Italic") + " and HiGreen"))
	fmt.Println(HiGreen("HiGreen and " + Underlined("Underlined") + " and HiGreen"))
	fmt.Println(HiGreen("HiGreen and " + Reversed("Reversed") + " and HiGreen"))
}

func TestHiYellow(t *testing.T) {
	fmt.Println(HiYellow("HiYellow and " + Bold("Bold") + " and HiYellow"))
	fmt.Println(HiYellow("HiYellow and " + Bright("Bright") + " and HiYellow"))
	fmt.Println(HiYellow("HiYellow and " + Faint("Faint") + " and HiYellow"))
	fmt.Println(HiYellow("HiYellow and " + Italic("Italic") + " and HiYellow"))
	fmt.Println(HiYellow("HiYellow and " + Underlined("Underlined") + " and HiYellow"))
	fmt.Println(HiYellow("HiYellow and " + Reversed("Reversed") + " and HiYellow"))
}

func TestHiBlue(t *testing.T) {
	fmt.Println(HiBlue("HiBlue and " + Bold("Bold") + " and HiBlue"))
	fmt.Println(HiBlue("HiBlue and " + Bright("Bright") + " and HiBlue"))
	fmt.Println(HiBlue("HiBlue and " + Faint("Faint") + " and HiBlue"))
	fmt.Println(HiBlue("HiBlue and " + Italic("Italic") + " and HiBlue"))
	fmt.Println(HiBlue("HiBlue and " + Underlined("Underlined") + " and HiBlue"))
	fmt.Println(HiBlue("HiBlue and " + Reversed("Reversed") + " and HiBlue"))
}

func TestHiMagenta(t *testing.T) {
	fmt.Println(HiMagenta("HiMagenta and " + Bold("Bold") + " and HiMagenta"))
	fmt.Println(HiMagenta("HiMagenta and " + Bright("Bright") + " and HiMagenta"))
	fmt.Println(HiMagenta("HiMagenta and " + Faint("Faint") + " and HiMagenta"))
	fmt.Println(HiMagenta("HiMagenta and " + Italic("Italic") + " and HiMagenta"))
	fmt.Println(HiMagenta("HiMagenta and " + Underlined("Underlined") + " and HiMagenta"))
	fmt.Println(HiMagenta("HiMagenta and " + Reversed("Reversed") + " and HiMagenta"))
}

func TestHiCyan(t *testing.T) {
	fmt.Println(HiCyan("HiCyan and " + Bold("Bold") + " and HiCyan"))
	fmt.Println(HiCyan("HiCyan and " + Bright("Bright") + " and HiCyan"))
	fmt.Println(HiCyan("HiCyan and " + Faint("Faint") + " and HiCyan"))
	fmt.Println(HiCyan("HiCyan and " + Italic("Italic") + " and HiCyan"))
	fmt.Println(HiCyan("HiCyan and " + Underlined("Underlined") + " and HiCyan"))
	fmt.Println(HiCyan("HiCyan and " + Reversed("Reversed") + " and HiCyan"))
}

func TestHiWhite(t *testing.T) {
	fmt.Println(HiWhite("HiWhite and " + Bold("Bold") + " and HiWhite"))
	fmt.Println(HiWhite("HiWhite and " + Bright("Bright") + " and HiWhite"))
	fmt.Println(HiWhite("HiWhite and " + Faint("Faint") + " and HiWhite"))
	fmt.Println(HiWhite("HiWhite and " + Italic("Italic") + " and HiWhite"))
	fmt.Println(HiWhite("HiWhite and " + Underlined("Underlined") + " and HiWhite"))
	fmt.Println(HiWhite("HiWhite and " + Reversed("Reversed") + " and HiWhite"))
}

func TestOrange(t *testing.T) {
	fmt.Println(Orange("Orange and " + Bold("Bold") + " and Orange"))
	fmt.Println(Orange("Orange and " + Bright("Bright") + " and Orange"))
	fmt.Println(Orange("Orange and " + Faint("Faint") + " and Orange"))
	fmt.Println(Orange("Orange and " + Italic("Italic") + " and Orange"))
	fmt.Println(Orange("Orange and " + Underlined("Underlined") + " and Orange"))
	fmt.Println(Orange("Orange and " + Reversed("Reversed") + " and Orange"))
}

func TestHiOrange(t *testing.T) {
	fmt.Println(HiOrange("HiOrange and " + Bold("Bold") + " and HiOrange"))
	fmt.Println(HiOrange("HiOrange and " + Bright("Bright") + " and HiOrange"))
	fmt.Println(HiOrange("HiOrange and " + Faint("Faint") + " and HiOrange"))
	fmt.Println(HiOrange("HiOrange and " + Italic("Italic") + " and HiOrange"))
	fmt.Println(HiOrange("HiOrange and " + Underlined("Underlined") + " and HiOrange"))
	fmt.Println(HiOrange("HiOrange and " + Reversed("Reversed") + " and HiOrange"))
}

func TestPink(t *testing.T) {
	fmt.Println(Pink("Pink and " + Bold("Bold") + " and Pink"))
	fmt.Println(Pink("Pink and " + Bright("Bright") + " and Pink"))
	fmt.Println(Pink("Pink and " + Faint("Faint") + " and Pink"))
	fmt.Println(Pink("Pink and " + Italic("Italic") + " and Pink"))
	fmt.Println(Pink("Pink and " + Underlined("Underlined") + " and Pink"))
	fmt.Println(Pink("Pink and " + Reversed("Reversed") + " and Pink"))
}

func TestHiPink(t *testing.T) {
	fmt.Println(HiPink("HiPink and " + Bold("Bold") + " and HiPink"))
	fmt.Println(HiPink("HiPink and " + Bright("Bright") + " and HiPink"))
	fmt.Println(HiPink("HiPink and " + Faint("Faint") + " and HiPink"))
	fmt.Println(HiPink("HiPink and " + Italic("Italic") + " and HiPink"))
	fmt.Println(HiPink("HiPink and " + Underlined("Underlined") + " and HiPink"))
	fmt.Println(HiPink("HiPink and " + Reversed("Reversed") + " and HiPink"))
}

func TestConvFuncs(t *testing.T) {
	PrintBlack("Black")
	PrintBgBlack("Black")
	PrintRed("Red")
	PrintBgRed("Red")
	PrintGreen("Green")
	PrintBgGreen("Green")
	PrintYellow("Yellow")
	PrintBgYellow("Yellow")
	PrintBlue("Blue")
	PrintBgBlue("Blue")
	PrintMagenta("Magenta")
	PrintBgMagenta("Magenta")
	PrintCyan("Cyan")
	PrintBgCyan("Cyan")
	PrintWhite("White")
	PrintBgWhite("White")
	PrintHiBlack("HiBlack")
	PrintBgHiBlack("HiBlack")
	PrintHiRed("HiRed")
	PrintBgHiRed("HiRed")
	PrintHiGreen("HiGreen")
	PrintBgHiGreen("HiGreen")
	PrintHiYellow("HiYellow")
	PrintBgHiYellow("HiYellow")
	PrintHiBlue("HiBlue")
	PrintBgHiBlue("HiBlue")
	PrintHiMagenta("HiMagenta")
	PrintBgHiMagenta("HiMagenta")
	PrintHiCyan("HiCyan")
	PrintBgHiCyan("HiCyan")
	PrintHiWhite("HiWhite")
	PrintBgHiWhite("HiWhite")
	PrintOrange("Orange")
	PrintBgOrange("Orange")
	PrintHiOrange("HiOrange")
	PrintBgHiOrange("HiOrange")
	PrintPink("Pink")
	PrintBgPink("Pink")
	PrintHiPink("HiPink")
	PrintBgHiPink("HiPink")
}

// Prints out if the enviroment variable NO_COLOR is set
func TestNO_COLOR(t *testing.T) {
	_, noColor := os.LookupEnv("NO_COLOR")
	if noColor {
		fmt.Println("No color")
	} else {
		fmt.Println("color")
	}
}

func TestColor(t *testing.T) {
	type args struct {
		c   Colr
		str string
	}
	tests := []struct {
		name  string
		color Colr
		text  string
		want  string
	}{
		{"black", ColBlack, "HELLO", "\x1b[38;5;0mHELLO\x1b[39m"},
		{"yellow", ColYellow, "HELLO", "\x1b[38;5;3mHELLO\x1b[39m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Color(tt.color, tt.text); got != tt.want {
				t.Errorf("Color() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBgColor(t *testing.T) {
	type args struct {
		c   Colr
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BgColor(tt.args.c, tt.args.str); got != tt.want {
				t.Errorf("BgColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
