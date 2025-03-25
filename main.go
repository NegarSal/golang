package main

import "fmt"

func main() {

  //Array
  var ages1 [3]int = [3]int{20, 25, 30}
  var ages2 = [3]int{20, 25, 30}
  names := [4]string{"yoshi", "mario", "peach", "bowser"}

  fmt.Println(ages1, ages2)
  fmt.Println(ages1, len(ages1))
  fmt.Println(names, len(names))

}