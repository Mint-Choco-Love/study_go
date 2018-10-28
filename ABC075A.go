package main

import (
  "fmt"
  "sort"
)
func main() {
  var vec []int
  for i := 0; i < 3; i++ {
    var temp int
    fmt.Scanf("%d", &temp)
    vec = append(vec, temp)
  }
  sort.Ints(vec)
  var ret int
  if vec[0] == vec[1] {
    ret = vec[2]
  } else {
    ret = vec[0]
  }

  fmt.Println(ret)
}
