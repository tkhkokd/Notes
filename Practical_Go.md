# Practical Go

## Reading data file

Reading 2 columns data sesparated by comma and print it out.

```go
package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)

func main() {
  xys, err := readData("DataScience/data.txt")
  if err != nil {
    log.Fatalf("could not read data.txt: %v", err)
  }
  for _, xy := range xys {
    fmt.Println(xy.x, xy.y)
  }
}

type xy struct { x, y float64 }

func readData(path string) ([]xy, error){
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  var xys []xy
  s := bufio.NewScanner(f)
  for s.Scan() {
      var x, y float64
      _, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y) // only err is needed to be assigned
      if err != nil {
        log.Printf("Discarding bad data %q: %v", s.Text(), err)
      }
      xys = append(xys, xy{x, y})
  }
  if err := s.Err(); err != nil {
    return nil, fmt.Errorf("could not scan: %v")
  }
  return xys, nil
}
```

## Parsing JSON
1. Use ```encoding/json``` package.
2. Define structs corresponding to the JSON object structure.
3. Unmarshal the JSON object to the structs.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Country struct {
    Name string    `json:"name"`    // `json:xxx` is called tags (Optional)
    States []State `json:"states"`  // Tags get the corresponding part of the json obj
}

type State struct {
    Name string    `json:"name"`
    Capital string `json:"capital"`
    Population int `json:"population"`
}

func main() {
    jsonStr := `
{
  "name": "The United States",
  "States": [
    {
      "name": "Maryland",
      "capital": "Annapolis",
      "population": 38394
    },
    {
      "name": "Ilinois",
      "capital": "Chicago",
      "population": 2707120
    },
    {
      "name": "New York",
      "capital": "Albany",
      "population": 97856
    }
  ]
}
`
    jsonBytes := ([]byte)(jsonStr)
    data := new(Country)

    if err := json.Unmarshal(jsonBytes, data); err != nil {   // Unmarshall maps the json obj to the struct
        fmt.Println("JSON Unmarshal error:", err)
        return
    }

    fmt.Println(data.Name)
    fmt.Println(data.States[0].Name)
    fmt.Println(data.States[1].Capital)
    fmt.Println(data.States[2].Population)
}
```
## Sorting strings

```go
package main

import "fmt"
import "sort"

func main() {
  strs := []string{"c", "a", "b"}
  sort.Strings(strs)
  fmt.Println("Strings: ", strs)

  ints := []int{7, 2, 4}
  sort.Ints(ints)
  fmt.Println("Ints: ", ints)

  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted:" s)
}

=>
Strings:  [a b c]
Ints:  [2 4 7]
Sorted: true
```

## Type Switches

```.(type)```separates type info from the variable.

```go
switch v := i.(type) {
case int:
    // here v has type T
case string:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
