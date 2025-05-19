package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	m := &name

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}
