package main

import (
  "fmt"
)

func main() {
  var N, A, B int
  fmt.Scanf("%d %d %d", &N, &A, &B)

  if N * A < B {
    fmt.Println(N * A)
  } else {
    fmt.Println(B)
  }
}
