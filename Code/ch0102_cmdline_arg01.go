package main

import (
  "fmt"
  "os"
)

func main() {
  var s, sep string
  // declares two vars
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}


// ex. go run ch02_cmdline_arg.go javathehat
