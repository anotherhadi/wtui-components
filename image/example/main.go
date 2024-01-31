package main

import (
	"github.com/anotherhadi/wtui-components/image"
)

func main() {
	// Change the path with your image
	err := image.Image("/home/hadi/Pictures/tokyo.png", 66, 20)
	if err != nil {
		panic(err)
	}
}
