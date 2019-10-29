package color256

import (
	"fmt"
	"testing"
)

func TestManual(t *testing.T) {
	Init()
	Random("bla")
	Random("Random Color %s", "wow")
	Red("Red")
	fmt.Print(Red(Bold("Red Bold")))

	fmt.Println()

}
