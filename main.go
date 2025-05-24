package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
