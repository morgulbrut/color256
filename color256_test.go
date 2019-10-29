package color256

import (
	"fmt"
	"testing"
)

func TestManual(t *testing.T) {
	Init()
	fmt.Println(Red("Red") + " and " + Green("Green"))
	fmt.Println(Red(Italic("Red Italic")))
	fmt.Println(Red(Bold("Red Bold")))
	fmt.Println(Red(Bold(Italic("Red Bold Italic"))))
	fmt.Println(Red(Reversed("Red Bold Reversed")))

	fmt.Println(Orange("Orange") + " and " + Green("Green"))
	fmt.Println(Orange(Italic("Orange Italic")))
	fmt.Println(Orange(Bold("Orange Bold")))
	fmt.Println(Orange(Bold(Italic("Orange Bold Italic"))))
	fmt.Println(Orange(Reversed("Orange Bold Reversed")))

	fmt.Println(HiOrange("HiOrange") + " and " + Green("Green"))
	fmt.Println(HiOrange(Italic("HiOrange Italic")))
	fmt.Println(HiOrange(Bold("HiOrange Bold")))
	fmt.Println(HiOrange(Bold(Italic("HiOrange Bold Italic"))))
	fmt.Println(HiOrange(Reversed("HiOrange Bold Reversed")))

	fmt.Println(Yellow("Yellow") + " and " + Green("Green"))
	fmt.Println(Yellow(Italic("Yellow Italic")))
	fmt.Println(Yellow(Bold("Yellow Bold")))
	fmt.Println(Yellow(Bold(Italic("Yellow Bold Italic"))))
	fmt.Println(Yellow(Reversed("Yellow Bold Reversed")))

	fmt.Println(HiYellow("HiYellow") + " and " + Green("Green"))
	fmt.Println(HiYellow(Italic("HiYellow Italic")))
	fmt.Println(HiYellow(Bold("HiYellow Bold")))
	fmt.Println(HiYellow(Bold(Italic("HiYellow Bold Italic"))))
	fmt.Println(HiYellow(Reversed("HiYellow Bold Reversed")))

	fmt.Println()

}
