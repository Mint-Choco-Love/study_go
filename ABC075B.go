package main

import (
  "fmt"
)
var H, W int

var dx = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
var dy = [8]int{1, 1, 0, -1, -1, -1, 0, 1}

var grid []string
var ret [51][51]int

func main() {
  fmt.Scanf("%d %d", &H, &W)
  for i := 0; i < H; i++ {
    var temp string
    fmt.Scanf("%s", &temp)
    grid = append(grid, temp)
  }

  for i := 0; i < H; i++ {
    for j := 0; j < W; j++ {
      for k := 0; k < 8; k++ {
        if grid[i][j] != '#' {
          continue
        }

        nx, ny := i + dx[k], j + dy[k]
        if nx < 0 || ny < 0 || nx >= H || ny >= W {
          continue
        }
        ret[nx][ny] += 1
      }
    }
  }

  for i := 0; i < H; i++ {
    for j := 0; j < W; j++ {
      if grid[i][j] == '#' {
        fmt.Print("#")
      } else {
        fmt.Print(ret[i][j])
      }
    }
    fmt.Println()
  }

}
