package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]int)
  //make initializes a map of a given type, which is int
  //map holds a set of key/value pair, which can be stored, retrieved, tested
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
  // a typical use of bufio.NewScanner(os.Stdin)
    counts[input.Text()]++
    // equivalent to line := input.Text()
    //               counts[line] = counts[line] + 1
    fmt.Println(counts)
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
