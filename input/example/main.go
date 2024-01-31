package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/input"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	userInput, _ := input.Input("What's your name?", "Type your name")
	fmt.Println(userInput)

	// With options
	opts := wtopts.DefaultOpts()
	opts.AccentColor = 4
	userInput, _ = input.Input("What's your name?", "Type your name")
	fmt.Println(userInput)
}
