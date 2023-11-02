package main

import (
  "fmt"
  alphabet "github.com/kekcleader/go-alphabet"
)

func main() {
  fmt.Println("Sorter")
  m := []string{"Museum", "Hello", "World", "Apple", "Win"}
  alphabet.Sort(m)
  for _, x := range m {
    fmt.Println(">", x)
  }
  fmt.Println(" (end of list)")
}
