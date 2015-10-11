package main

import (
  "net"
  "fmt"
  "log"
)

func main() {
  host := "google.com"
  addrs, err := net.LookupHost(host)
  if err == nil {
    fmt.Printf("Addresses for %s are\n", host)
    for _, addr := range addrs {
      names, err := net.LookupAddr(addr)
      if err == nil {
        for _, name := range names {
          fmt.Printf("%s, Reverse lookup for address %s is %s\n\n", addr, addr, name)
        }
      } else {
        fmt.Printf("Error doing a reverse lookup for address %s, %v\n", addr, err)
      }
    }
  } else {
    log.Fatalf("Error on LookupHost: %v\n", err)
  }

  fmt.Println("--------------------------------------------------------")
  cname, err := net.LookupCNAME(host)
  if err == nil {
    fmt.Printf("CNAME for %s is %s", host, cname)
  } else {
    log.Printf("Error on LookupCNAME: %v\n", err)
  }

 fmt.Println("--------------------------------------------------------")
  mxs, err := net.LookupMX(host)
  if err == nil {
    fmt.Printf("Printing MX records for %s\n", host)
    for _, mx := range mxs {
      fmt.Printf("Host %s, Pref %d\n", mx.Host, mx.Pref)
    }
  } else {
    log.Printf("Error on LookupMX: %v\n", err)
  }

fmt.Println("--------------------------------------------------------")
 txts, err := net.LookupTXT(host)
 if err == nil {
   fmt.Printf("Printing DNS text records for %s\n", host)
   for _, txt := range txts {
     fmt.Printf("%s\n", txt)
   }
 } else {
   fmt.Printf("Error on LookupTXT for %s, %v\n", host, err)
}

fmt.Println("--------------------------------------------------------")
 nss, err := net.LookupNS(host)
 if err == nil {
   fmt.Printf("Printing NS records for %s\n", host)
   for _, ns := range nss {
     fmt.Printf("%s\n", ns.Host)
   }
 } else {
   fmt.Printf("Error on LookupNS for %s, %v\n", host, err)
 }
}
