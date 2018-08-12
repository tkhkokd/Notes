package main

import (
  "fmt"
  "reflect"
)

func main() {
  c := test()
  fmt.Println(reflect.TypeOf(c))
}

func test() *int {
  v := 1
  return &v
}
