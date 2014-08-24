package main

import "fmt"

func plus(a, b int) int {
  return a + b
}

func vals() (int, int) {
  return 3, 7
}

func sum(nums ... int) {
  fmt.Println(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }

  fmt.Println(total)
}

func intSeq() func() int {
    i := 0
    return func() int {
      i += 1
      return i
    }
}

func main() {
  res := plus(1, 2)
  fmt.Println("1+2 =", res)

  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)

  sum(1)
  sum(1, 2)
  sum(1, 2, 3)

  newInts := intSeq()
  fmt.Println(newInts())
}
