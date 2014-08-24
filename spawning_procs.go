package main

import (
  "fmt"
  "os/exec"
  "io/ioutil"
  )


func main() {
  calCmd := exec.Command("cal", "-y", "2014")
  calOut, err := calCmd.Output()
  if(err != nil) {
    panic(err)
  } else {
    fmt.Println(string(calOut))
  }

  teeCmd := exec.Command("tee")

  teeIn, _ := teeCmd.StdinPipe()
  teeOut, _ := teeCmd.StdoutPipe()
  teeCmd.Start()
  teeIn.Write([]byte("hello 123 \n foo bar goo"))
  teeIn.Close()
  teeBytes, _ := ioutil.ReadAll(teeOut)
  teeCmd.Wait()
  fmt.Println(string(teeBytes))
}
