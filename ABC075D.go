package main

import (
  "fmt"
  "math"
//  "sort"
)

var N, K int

type Pos struct {
  X int64
  Y int64
}

var vec []Pos

func abs(a int64) int64 {
  if a >= 0 {
    return a
  } else {
    return -a
  }
}

func max(a, b int64) int64 {
  if a > b {
    return a
  } else {
    return b
  }
}

func min (a, b int64) int64 {
  if a < b {
    return a
  } else {
    return b
  }
}

func test(lt, rd Pos) int {
  ret := 0
  for i := 0; i < N; i++ {
    cur := vec[i]
    if(lt.X <= cur.X && cur.X <= rd.X && rd.Y <= cur.Y && cur.Y <= lt.Y) {
      ret += 1
    }
  }
  return ret
}

func getArea(lt, rd Pos) int64 {
  var ret int64 = (rd.X - lt.X) * (lt.Y - rd.Y)
  return ret
}

func main() {
  fmt.Scanf("%d %d", &N, &K)

  for i := 0; i < N; i++ {
    var x, y int64
    fmt.Scanf("%d %d", &x, &y)
    vec = append(vec, Pos{x, y})
  }

  var ret int64 = math.MaxInt64


  for i := 0; i < N; i++ {
    for j := 0; j < N; j++ {
      for k := 0; k < N; k++ {
        for l := 0; l < N; l++ {
          tl := Pos{vec[i].X, vec[j].Y}
          tr := Pos{vec[k].X, vec[l].Y}
          if tl.X > tr.X {
            continue
          }
          if tl.Y < tr.Y {
            continue
          }
          retN := test(tl, tr)
          retArea := getArea(tl, tr)
          fmt.Printf("(%d,%d), (%d,%d) : %d with retN: %d\n", tl.X, tl.Y, tr.X, tr.Y, retArea, retN)
          if retN >= K && retArea < ret {
            ret = retArea
          }
        }
      }
    }
  }


  fmt.Println(ret)
}
//9223372036854775807
//3999999992000000004
