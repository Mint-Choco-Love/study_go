package main

import (
  "fmt"
)
var str, t string
var quest []int

func check(idx int) bool {
  if len(str) - idx < len(t) {
    return false
  }
  for i := idx; i < idx + len(t); i++{
    if str[i] != t[i-idx] && str[i] != '?' {
      return false
    }
  }
  return true
}

func main() {
  fmt.Scanf("%s\n%s", &str, &t)
  if str == t {
    fmt.Println(str)
    return
  }

  quest = make([]int, len(str))
  cnt := 0
  for i := len(str) - 1; i >= 0; i-- {
    if str[i] != '?' {
      cnt = 0
    } else {
      cnt += 1
      quest[i] = cnt
    }
  }
  idx := -1
  for i := 0; i < len(str); i++ {
    if(check(i)) {
      idx = i
    }
  }
  if idx == -1 {
    fmt.Println("UNRESTORABLE")
  } else {
    for i := 0; i < len(str); i++ {
      if i == idx {
        fmt.Print(t)
        i += len(t) - 1
      } else if str[i] == '?' {
        fmt.Print("a")
      } else {
        fmt.Print(string(str[i]))
      }
    }
    fmt.Println()
  }

}
