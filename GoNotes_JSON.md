# JSON in Go

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
