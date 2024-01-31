package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/ansi"
)

func main() {
	fmt.Print(ansi.CursorHome)
	fmt.Print(ansi.ScreenClear)
	fmt.Print(ansi.Red + "Hello ")
	fmt.Print(ansi.BgBlue + "World")
	fmt.Print(ansi.CursorMove(2, 2) + ansi.Reset + ansi.Fg256(99) + "I'm Hadi")
	fmt.Print(ansi.Reset)
}
