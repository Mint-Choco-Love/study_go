package main

import (
  "fmt"
  "bufio"
  "flag"
  "io"
  "os"
)

func lineByLine(file string) error {
  var err error

  f, err := os.Open(file)
  if err != nil {
    return err
  }
  defer f.Close()

  r := bufio.NewReader(f)
  for {
    line, err := r.ReadString('\n')
    if err == io.EOF{
      break
    } else if err != nil {
      fmt.Println("error reading file %s", err)
      break
    }
    fmt.Print(line)
  }
  return nil
}

func main() {
  flag.Parse()
  if len(flag.Args()) == 0 {
    fmt.Println("usage: byLine <file1> [<File2> ...]\n")
    return
  }

  for _, file := range flag.Args() {
    err := lineByLine(file)
    if err != nil {
      fmt.Println(err)
    }
  }
}
