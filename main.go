package main

import (
  "log"
  "net/http"
)

func main() {
  router := NewRouter()
  log.Fatal(http.ListenListenAndServe(":8080", router))
}
