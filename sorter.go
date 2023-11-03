package main

import (
  "fmt"
  alphabet "github.com/kekcleader/go-alphabet"
)

func main() {
  fmt.Println("Sorter")
  m := []string{"дорис", "клот", "иргиш", "ИЛЛИОТ",
                "дука", "груж", "СМАЙД", "арбуэ",
                "АЛЁНА", "АЛЕНА", "АЛЁ", "АЛЕ-ОП"}
  alphabet.SortFold(m)
  for _, x := range m {
    fmt.Println(">", x)
  }
  fmt.Println(" (end of list)")
}
