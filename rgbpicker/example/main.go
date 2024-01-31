package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/rgbpicker"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	rgb, _ := rgbpicker.RGBPicker()
	fmt.Println(rgb)

	// With options
	opts := wtopts.DefaultOpts()
	opts.DefaultBlue = 123
	opts.DefaultGreen = 23
	rgb, _ = rgbpicker.RGBPicker(opts)
	fmt.Println(rgb)
}
