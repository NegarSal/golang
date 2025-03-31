package main

import (
	"fmt"
	"strings"
)

func main() {
  greeting := "hello there friends!"

  fmt.Println(strings.Contains(greeting, "hello"))
  fmt.Println(strings.Contains(greeting, "hello!"))
  fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
 
  // The original value is unchanged
    fmt.Println("original string value =", greeting)

}