package main

import "net/http"
import "fmt"

func main() {
  
  var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, "Hello World")
  }
  server := http.Server{
    Addr: "localhost:8009",
    Handler: handler,
  }
  
  err := server.ListenAndServe()
  if err != nil {
    panic(err)
  }
}

