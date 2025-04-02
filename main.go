package main

import (
	"fmt"
)

func main() {

	names := []string{"mario", "luigi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	for index, value := range names{
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names{
		fmt.Printf("the value is %v \n", value)
	}

	for _, value := range names{
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}
	fmt.Println(names)
}
