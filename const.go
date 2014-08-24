package main

import (
  "fmt"
  "math"
)

const (
  Red = 1 << iota
  Blue
  Green
)

func main() {
var flag bool = false
fmt.Println(Red)
fmt.Println(Blue)
fmt.Println(Green)
fmt.Println(flag)
fmt.Println(math.MaxInt64)
fmt.Println(math.MaxInt32)
fmt.Println(math.MaxInt16)
fmt.Println(math.MaxInt8)
fmt.Println(uint64(math.MaxUint64))
fmt.Println(uint32(math.MaxUint32))
fmt.Println(uint16(math.MaxUint16))
fmt.Println(uint8(math.MaxUint8))
}
