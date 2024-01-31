package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/checkbox"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	checked, _ := checkbox.Checkbox([]string{
		"Option 1",
		"Option 2",
		"Option 3",
		"Option 4",
		"Option 5",
		"Option 6",
		"Option 7",
		"Option 8",
		"Option 9",
	})
	fmt.Println(checked)

	// With options
	opts := wtopts.DefaultOpts()
	opts.MaxRows = 6
	checked, _ = checkbox.Checkbox([]string{
		"Option 1",
		"Option 2",
		"Option 3",
		"Option 4",
		"Option 5",
		"Option 6",
		"Option 7",
		"Option 8",
		"Option 9",
	}, opts)
	fmt.Println(checked)
}
