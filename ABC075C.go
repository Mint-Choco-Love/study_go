package main

import (
  "fmt"
)
var N, M int

var adj[51][]int

func check(from, to int) bool{
  visited := make([]bool, N+1)

  cnt := 0
  for i := 1; i <= N; i++ {
    if visited[i] {
      continue
    }
    cnt += 1
    dfs(i, visited, from, to)
  }
  if cnt > 1 {
    return true
  } else {
    return false
  }
}

func dfs(cur int, visited []bool, from, to int) {
  visited[cur] = true
  for i := 0; i < len(adj[cur]); i++ {
    if cur == from && adj[cur][i] == to {
      continue
    }
    next := adj[cur][i]
    if visited[next] {
      continue
    }
    dfs(next, visited, from, to)
  }
}

func main() {
  fmt.Scanf("%d %d", &N, &M)

  for i := 0; i < M; i++ {
    var a, b int
    fmt.Scanf("%d %d", &a, &b)
    adj[a] = append(adj[a], b)
    adj[b] = append(adj[b], a)
  }
  ret := 0

  for i := 1; i <= N; i++ {
    for j := 0; j < len(adj[i]); j++ {
      if(check(i, adj[i][j])) {
        ret += 1
//        fmt.Printf("fronm %d to %d\n", i, adj[i][j])
      }
    }
  }
  fmt.Println(ret)
}
