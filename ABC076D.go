package main

import (
  "fmt"
)

func min(a, b float64) float64 {
  if a > b {
    return b
  } else {
    return a
  }
}

func get(time_interval, cur_v, next_v float64) float64 {
  var ret float64 = 0
  ret += next_v * time_interval

  if next_v > cur_v { // 속도가 늘어나는 경우
    offset := next_v - cur_v
    ret -= (offset * offset / 2)
  } else if next_v < cur_v { // 속도가 줄어드는 경우
    offset := cur_v - next_v
    ret -= (offset * offset / 2)
  }
  return ret
}

func main() {
  var t, v []float64
  var N int
  fmt.Scanf("%d", &N)

  for i := 0; i < N; i++ {
    var temp float64
    fmt.Scanf("%f", &temp)
    t = append(t, temp)
  }
  for i := 0; i < N; i++ {
    var temp float64
    fmt.Scanf("%f", &temp)
    v = append(v, temp)
  }


  var ret float64 = 0
  var cur_v float64 = 0
  var cur_t float64 = 0
  v[len(v)-1] = 0


  for i := 0; i < N; i++ {
    next_v := v[i]
    next_t := cur_t + t[i]

    fst := min(next_v - cur_v, next_t - cur_t)
    snd := next_t - fst
    ret += get(t[i], cur_v, fst)
    ret += get(t[i], cur_v, snd)

    fmt.Printf("cur_t: %f, cur_v: %f, next_t: %f, next_v: %f, fst: %f, snd: %f, ret: %f\n", cur_t, cur_v, next_t, next_v, fst, snd, ret)
    cur_v = next_v
    cur_t = next_t
  }


  fmt.Printf("%f\n", ret)
}
