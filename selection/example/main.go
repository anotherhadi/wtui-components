package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/selection"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	selected, _ := selection.Selection([]string{"One", "Two", "Three"})
	fmt.Println(selected)

	// With options
	opts := wtopts.DefaultOpts()
	opts.MaxRows = 4
	selected, _ = selection.Selection([]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven"}, opts)
	fmt.Println(selected)
}
