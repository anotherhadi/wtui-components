package main

import (
	"fmt"

	"github.com/anotherhadi/wtui-components/getsize"
)

func main() {
	cols, rows, _ := getsize.GetSize()
	fmt.Println("Cols:", cols, "Rows:", rows)
}
