package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func main() {
	// group A types -> strings, ints, booleans, floats, arrays, structs
	name := "tifa"

	updateName(name)

	fmt.Println(name)
}
