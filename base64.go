package main

import b64 "encoding/base64"
import "fmt"

func main() {
  data := "?nvacl_user=nvellanki"

  sEnc := b64.StdEncoding.EncodeToString([]byte(data))
  fmt.Println(sEnc)

  sDec, _ := b64.StdEncoding.DecodeString(sEnc)
  fmt.Println(string(sDec))
  uEnc := b64.URLEncoding.EncodeToString([]byte(data))
  fmt.Println(string(uEnc))

  uDec, _ := b64.URLEncoding.DecodeString(uEnc)
  fmt.Println(string(uDec))
}
