package main

import (
  "fmt"
  "net"
  "os"
)

func lookIP(address string) ([]string, error) {
  hosts, err := net.LookupAddr(address)
  if err != nil {
    return nil, err
  }
  return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
  IPs, err := net.LookupHost(hostname)
  if err != nil {
    return nil, err
  }
  return IPs, nil
}

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("Please provide an argument!")
    return
  }
  for i := 1; i < len(arguments); i++ {
    input := arguments[i]
    IPaddress := net.ParseIP(input)
    fmt.Printf("input is %s\n", input)


    if IPaddress == nil {
      IPs, err := lookHostname(input)
      if err == nil {
        for _, singleIP := range IPs {
          fmt.Println(singleIP)
        }
      }
    } else {
      hosts, err := lookIP(input)
      if err == nil {
        for _, hostname := range hosts {
          fmt.Println(hostname)
        }
      }
    }
    fmt.Println()
  }
}
