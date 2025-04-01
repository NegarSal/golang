package main

import (
	"fmt"
	"sort"
)

func main() {
  ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

  sort.Ints(ages)
  fmt.Println(ages)

  index := sort.SearchInts(ages, 30)
  fmt.Println(index)

  index2 := sort.SearchInts(ages, 90)
  fmt.Println(index2)

  names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
  sort.Strings(names)
  fmt.Println(names)
  
  fmt.Println(sort.SearchStrings(names, "bowser"))
}