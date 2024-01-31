package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/floatingwindow"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	mybox := floatingwindow.FloatingWindow(15, 5, "My Title")
	for _, line := range mybox {
		fmt.Println(line)
	}

	// With options and no title
	opts := wtopts.DefaultOpts()
	opts.Style = "double"
	mybox = floatingwindow.FloatingWindow(20, 5, "", opts)
	for _, line := range mybox {
		fmt.Println(line)
	}
}
