package main

import (
	"github.com/anotherhadi/wtui-components/table"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	content := [][]string{
		{"Name", "Description", "Usage"},
		{"EditFirst", "Edit the first variable", "vlafzfa"},
		{"EditN", "Edit the N variable", "vlafzfa"},
	}

	table.Table(content)

	// With Options
	opts := wtopts.DefaultOpts()
	opts.FirstRowAccentColor = false
	opts.DefaultColLength = 40
	opts.Spacer = false

	table.Table(content, opts)
}
