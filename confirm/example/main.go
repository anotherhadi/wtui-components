package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/confirm"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	result, _ := confirm.Confirm()
	fmt.Println(result)

	// With options
	opts := wtopts.DefaultOpts()
	opts.Loop = false
	result, _ = confirm.Confirm(opts)
	fmt.Println(result)
}
