package main

import (
  "fmt"
)
func main() {
  var N int
  fmt.Scanf("%d", &N)
  var stk []int

  for i := 0; i < N; i++ {
    var str string
    fmt.Scanf("%s", &str)
    if str == "push" {
      var input int
      fmt.Scanf("%d", &input)
      stk = append(stk, input)

    } else if str == "top" {
      if len(stk) == 0 {
        fmt.Println(-1)
      } else {
        fmt.Println(stk[len(stk)-1])
      }

    } else if str == "size" {
      fmt.Println(len(stk))

    } else if str == "pop" {
      if len(stk) == 0 {
        fmt.Println(-1)
      } else {
        fmt.Println(stk[len(stk) - 1])
        stk = stk[:len(stk)-1]
      }

    } else if str == "empty" {
      if len(stk) > 0 {
        fmt.Println(0)
      } else {
        fmt.Println(1)
      }
    }
  }

}
