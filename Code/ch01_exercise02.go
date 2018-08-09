package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  s, sep := "", " "
  for i, arg := range os.Args[1:] {
    s = strconv.Itoa(i) + sep + arg
    fmt.Println(s)
  }
}

// convert int using Itoa func in strconv package
