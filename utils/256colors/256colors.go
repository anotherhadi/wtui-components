package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/ansi"
)

func main() {
	fmt.Println("256 Terminal Colors (Ordered):")

	// Standard colors
	fmt.Println("Standard Colors:")
	for i := 0; i < 16; i++ {
		fmt.Print(ansi.Bg256(uint8(i))+fmt.Sprintf("%3d", i)+ansi.Reset, " ")
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}

	// Extended colors
	fmt.Println("\nExtended Colors:")
	for r := 0; r < 6; r++ {
		for g := 0; g < 6; g++ {
			for b := 0; b < 6; b++ {
				color := 16 + r*36 + g*6 + b
				fmt.Print(ansi.Bg256(uint8(color))+fmt.Sprintf("%3d", color)+ansi.Reset, " ")
				if (color-15)%36 == 0 {
					fmt.Println()
				}
			}
		}
	}

	// Grayscale colors
	fmt.Println("\nGrayscale Colors:")
	for i := 232; i < 256; i++ {
		fmt.Print(ansi.Bg256(uint8(i))+fmt.Sprintf("%3d", i)+ansi.Reset, " ")
	}
}
