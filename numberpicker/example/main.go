package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/numberpicker"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	// By default only integers
	number, _ := numberpicker.NumberPicker()
	fmt.Println(number)

	// With options
	opts := wtopts.DefaultOpts()
	opts.Maximum = 1
	opts.Minimum = 0
	opts.Decimals = 1
	opts.Increment = 0.1

	number, _ = numberpicker.NumberPicker(opts)
	fmt.Println(number)
}
