package main


import (
  "fmt"
  "strings"
  )

func Index(vs *[]string, t string) int {
    for i, v := range *vs {
      if v == t {
        return i
      }
    }
    return -1
}

func Include(vs *[]string, t string) bool {
  return Index(vs, t) >= 0
}

func Any(vs *[]string, f func(string) bool) bool {
  for _, v := range *vs {
    if f(v) {
      return true
    }
  }
  return false
}

func Filter(vs *[]string, f func(string) bool) []string {
  temp := make([]string, 0)
  for _, v := range *vs {
    if f(v) {
      temp = append(temp, v)
    }
  }
  return temp
}

func Map(vs *[]string, f func(string) string) []string {
    temp := make([]string , len(*vs))
    for i, v := range *vs {
      temp[i] = f(v)
    }

    return temp
}

func main() {
  var strs = []string{"peach", "apple", "pear", "plum"}

  fmt.Println("Index of peach is ", Index(&strs, "peach"))
  fmt.Println("Index of apple is ", Index(&strs, "apple"))
  fmt.Println("Index of pear is ", Index(&strs, "pear"))
  fmt.Println("Index of plum is ", Index(&strs, "plum"))

  fmt.Println("Does slice include plum", Include(&strs, "plum"))
  fmt.Println("Does slice include peach", Include(&strs, "peach"))
  fmt.Println("Does slice include apple", Include(&strs, "apple"))
  fmt.Println("Does slice include peac", Include(&strs, "peac"))

  fmt.Println("Does any entry in slice has prefix pe: ", Any(&strs, func(v string) bool {
    return strings.HasPrefix(v, "pe")
    }))

  fmt.Println("Does any entry in slice has prefix re: ", Any(&strs, func(v string) bool {
    return strings.HasPrefix(v, "re")
    }))

  fmt.Println("Does any string in slice contain e: ", Filter(&strs, func(v string) bool {
    return strings.Contains(v, "e")
    }))

  fmt.Println("Does any string in slice contain v: ", Filter(&strs, func(v string) bool {
    return strings.Contains(v, "v")
    }))

  fmt.Println(Map(&strs, strings.ToUpper))
}
