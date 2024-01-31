package main

import "github.com/anotherhadi/wtui-components/shortcuts"

func main() {
	shortcuts.Shortcuts([][2]string{
		{"Ctrl+C", "Exit"},
		{"Ctrl+S", "Save"},
		{"P", "Previous Page"},
		{"N", "Next Page"},
	})
}
