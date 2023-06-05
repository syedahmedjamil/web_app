package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from Go App - UnChanged")
  })
  err := http.ListenAndServe("0.0.0.0:1026", nil)
  if err != nil {
    panic(err)
  } 
}
