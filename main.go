package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "wedge"
	return x
}

func main() {
	// group A types -> strings, ints, booleans, floats, arrays, structs
	name := "tifa"

	name = updateName(name)

	fmt.Println(name)
}
