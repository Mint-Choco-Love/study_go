package main

import (
  "fmt"
  "strconv"
)

func HarshadNumber(s string) int {
  ret := 0
  for i := 0; i < len(s); i++ {
    num, _ := strconv.Atoi(string(s[i]))
    ret += num
  }
  return ret
}

func main() {
  var N string
  fmt.Scanf("%s", &N)

  ret := HarshadNumber(N)
  NN, _ := strconv.Atoi(N)

  if NN % ret == 0 {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
