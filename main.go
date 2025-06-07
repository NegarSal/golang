package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
}

func main() {
	mybill := createBill()
}
