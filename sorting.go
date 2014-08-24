package main

import "fmt"
import "sort"

type ByLength []string

func (s ByLength) Len() int {
  return len(s)
}

func (s ByLength) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
  return len(s[i]) < len(s[j])
}

func main() {
  strs := []string{"c", "a", "b"}
  sort.Strings(strs)
  fmt.Println("strings: ", strs)

  ints := []int{50, 1, 4, 3}
  sort.Ints(ints)
  fmt.Println(ints)
  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted:", s)

  fruits := []string{"peach", "banana", "kiwi"}
  sort.Sort(ByLength(fruits))
  fmt.Println(fruits)
}
