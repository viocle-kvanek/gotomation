// +build ignore

package main

import (
	"fmt"

	"github.com/viocle-kvanek/gotomation"
)

func main() {
	screen, _ := gotomation.GetMainScreen()
	x, y := screen.Mouse().GetPosition()
	fmt.Printf("Mouse Position: %d, %d\n", x, y)
}
