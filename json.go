package main

import (
  "fmt"
  "encoding/json"
)

type Box struct {
  Fruits []string
  NumFruits int
}

func main() {
  boolB, _ := json.Marshal(true)
  fmt.Println(string(boolB))

  intB, _ := json.Marshal(1)
  fmt.Println(string(intB))

  strB, _ := json.Marshal("gopher")
  fmt.Println(string(strB))

  slc := []string{"apple", "peach", "pear"}
  sliceB, _ := json.Marshal(slc)
  fmt.Println(string(sliceB))

  maps := map[string]int{"apple": 1, "lettuce": 2}
  mapB, _ := json.Marshal(maps)
  fmt.Println(string(mapB))

  box := Box{Fruits: []string{"apple", "peach", "pear"}, NumFruits: 3}
  boxB, _ := json.Marshal(box)
  fmt.Println(string(boxB))

  str := `{"Fruits":["apple", "peach", "pear"],"NumFruits":3}`
  box = Box{}
  json.Unmarshal([]byte(str), &box)
  fmt.Printf("%#v\n", box)
}
