package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/list"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	result, _ := list.List([][2]string{
		{"First Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Second Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Another Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Other Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Test Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Last Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
	})
	fmt.Println(result)

	// With options

	opts := wtopts.DefaultOpts()
	opts.MaxCols = 30
	opts.MaxRows = 15
	result, _ = list.List([][2]string{
		{"First Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Second Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Another Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Other Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Test Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{".. Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		{"Last Option", "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
	}, opts)
	fmt.Println(result)
}
