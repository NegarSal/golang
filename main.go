package main

import "fmt"

var someName = "hello"

func main() {

	//strings

	//First way
	var nameOne string = "mario"

	//Second way
	var nameTwo="luigi"

	//Third way
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	//Update value

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	//4th way
	nameFour := "yoshi"
	fmt.Println(nameFour)

}
