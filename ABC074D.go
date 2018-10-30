package main

import (
  "fmt"
  "math"
)

func main() {
  const MAX = math.MaxInt32
  var N int
  fmt.Scanf("%d", &N)

  var dist [301][301]int
  var input [301][301]int

  for i := 0; i < N; i++ {
    for j := 0; j < N; j++ {
      fmt.Scanf("%d", &input[i][j])
    }
  }

  for k := 0; k < N; k++ {
    for i := 0; i < N; i++ {
      for j := 0; j < N; j++ {
        if i == j || j == k || k == i {
          continue
        }
        if input[i][j] > input[i][k] + input[k][j] {
          fmt.Println(-1)
          return
        } else if input[i][j] == input[i][k] + input[k][j] {
          dist[i][j] = MAX
        }
      }
    }
  }

  ret := 0
  for i := 0; i < N; i++ {
    for j := 0; j < N; j++ {
      if dist[i][j] == MAX {
        continue
      }
      ret += input[i][j]
    }
  }
  fmt.Println(ret / 2)
}
