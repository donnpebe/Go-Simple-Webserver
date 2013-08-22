package main

import (
  "fmt"
  "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello from Outer Space!")
}

func main() {
  var h Hello
  http.ListenAndServe("localhost:23666", h)
}