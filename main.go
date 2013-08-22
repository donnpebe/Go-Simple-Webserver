package main

import (
  "fmt"
  "net/http"
)

type String string
type Struct struct {
  Greeting string
  Punct string
  Who string
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, h)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, s.Who + " says" + s.Punct + " " + s.Greeting)
}

func main() {
  h := String("Hello Girls!")
  s := &Struct{
    "Hellow Girls!",
    ":",
    "Norman",
  }


  http.Handle("/",String("Welcome Girls..."))
  http.Handle("/string", h)
  http.Handle("/struct", s)
  http.ListenAndServe("localhost:23232",nil)
}