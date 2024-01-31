package spacer

import (
	"fmt"
)

func Draw(height int) []string {
	var blankline []string

	for i := 0; i < height; i++ {
		blankline = append(blankline, "")
	}

	return blankline
}

func Spacer(height int) {
	strs := Draw(height)
	for _, str := range strs {
		fmt.Println(str)
	}
}
