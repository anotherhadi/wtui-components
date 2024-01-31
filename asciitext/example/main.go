package main

import (
	"github.com/anotherhadi/wtui-components/asciitext"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	asciitext.Asciitext("Hello world")

	// With opts
	opts := wtopts.DefaultOpts()
	opts.LeftPadding = 9
	asciitext.Asciitext("I'm Hadi!", opts)
}
