package color256

import (
	"fmt"
	"testing"
)

func TestManual(t *testing.T) {
	Init()
	fmt.Println(StringRandom("bla"))
	fmt.Println(StringRandom("bla"))
	fmt.Println(StringRandom("bla"))
	fmt.Println(StringRandom("bla"))
	fmt.Println(StringRandom("bla"))

	fmt.Println(String(Cyan, "CYAN"))
	fmt.Println(String(Green, "GREEN"))
	fmt.Println(String(Orange, "ORANGE"))
	fmt.Println(StringFormat(Red, "RED", Faint))
	fmt.Println(StringFormatBoth(Red, Green, "RED", Faint))
	fmt.Println(StringFormatBoth(Red, Green, "RED", Underline))
	fmt.Println(StringBoth(Orange, Cyan, "ORANGE"))

	List()

	fmt.Println()
}
