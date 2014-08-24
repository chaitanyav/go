package hello

import (
  "fmt"
  "net/http"
  "appengine"
  "appengine/user"
  )

func init() {
  http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  context := appengine.NewContext(r)
  u := user.Current(context)
  if(u == nil) {
    url, error := user.LoginURL(context, r.URL.String())
    if error != nil {
      http.Error(w, error.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return
  }
  fmt.Fprintf(w, "Hello, %v!", u)
}
