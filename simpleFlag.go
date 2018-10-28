package main

import (
  "fmt"
  "flag"
)

func main() {
  minusK := flag.Bool("k", true, "k")
  minusO := flag.Int("O", 1, "O")
  flag.Parse()

  valueK := *minusK
  valueO := *minusO
  valueO += 1

  fmt.Println("-k:", valueK)
  fmt.Println("-O:", valueO)
}
