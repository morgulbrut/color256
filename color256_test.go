package color256

import (
	"fmt"
	"testing"
)

func TestManual(t *testing.T) {
	Init()
	Random("bla")
	Random("Random Color %s", "wow")
	fmt.Println(String(ColCyan, "CYAN"))
	fmt.Println(String(ColGreen, "GREEN"))
	fmt.Println(String(ColOrange, "ORANGE"))
	fmt.Println(StringFormat(ColRed, "RED", Faint))
	fmt.Println(StringFormatBoth(ColRed, ColGreen, "RED", Faint))
	fmt.Println(StringFormatBoth(ColRed, ColGreen, "RED", Underlined))
	fmt.Println(StringFormat(ColBlue, "Reverse Blue", Reversed))
	fmt.Println(StringBoth(ColOrange, ColCyan, "ORANGE"))
	//List()
	Black("Hello")
	Orange("Orange")
	HiRed("Hello %s", "World")
	HiRed("Hello %s", StringBlue("World"))
	WhiteReversed("Hello %s", "Reversed Colors")
	OrangeBold("Hello %s", "orange")
	GreenFaint("Hello %s", "green")
	fmt.Printf("Red: %s", StringRed("%s", "RED"))
	fmt.Println()
}
