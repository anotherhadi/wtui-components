package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/combobox"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	selected, _ := combobox.Combobox([]string{"One", "Two", "Three", "Four"})
	fmt.Println(selected)

	// With options
	opts := wtopts.DefaultOpts()
	opts.CaseSensitive = true
	selected, _ = combobox.Combobox([]string{"One", "Two", "Three", "Four"}, opts)
	fmt.Println(selected)
}
