package main

import (
  "fmt"
  "strings"
  "os"
)

func main() {
  fmt.Println(strings.Join(os.Args[1:], ","))
}

// func Join(a []string, separator string) string
