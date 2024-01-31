package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/rgbapicker"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	rgba, _ := rgbapicker.RGBAPicker()
	fmt.Println(rgba)

	// With options
	opts := wtopts.DefaultOpts()
	opts.DefaultRed = 200
	rgba, _ = rgbapicker.RGBAPicker(opts)
	fmt.Println(rgba)
}
