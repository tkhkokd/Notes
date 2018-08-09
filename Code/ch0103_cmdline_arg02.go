package main

import (
  "fmt"
  "os"
)

func main() {
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    // range produces index-value pair, _ => index, arg => value
    s += sep + arg
    sep = " "
    fmt.Println(arg)
  }
  fmt.Println(s)
}
