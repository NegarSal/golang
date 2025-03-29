package main

import "fmt"

func main() {

  //Array
  var ages1 [3]int = [3]int{20, 25, 30}
  var ages2 = [3]int{20, 25, 30}
  names := [4]string{"yoshi", "mario", "peach", "bowser"}
  names[1] = "luigi"

  fmt.Println(ages1, ages2)
  fmt.Println(ages1, len(ages1))
  fmt.Println(names, len(names))

  // slices (use arrays under the hood) 
  var scores = []int{100, 50, 60}
  scores[2] = 25
  scores = append(scores, 85)

  fmt.Println(scores, len(scores))

}