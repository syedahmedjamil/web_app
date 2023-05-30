package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from lesson !!")
  })
  err := http.ListenAndServe(":1026", nil)
  if err != nil {
    panic(err)
  } 
}
