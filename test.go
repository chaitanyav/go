package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println(1, "2", 2)
  var foo = map[int]string{}
  foo[1] = "foo"
  foo[2] = "bar"
  foo[3] = "goo"
  fmt.Println(foo[1])
  fmt.Println(foo[2])
  fmt.Println(foo[3])
  delete(foo, 3)
  value, ok := foo[3]
  fmt.Println(ok)
  fmt.Println(value)

  value, ok = foo[2]
  fmt.Println(ok)
  fmt.Println(value)
  var now = time.Now()
  fmt.Println(now.Clock())
}
