package main

import (
  "net/http"
  "os"
  "fmt"
 )

func main() {
preport := os.Args[1:]
fmt.Println("You can specify port as argument")
  if len(preport) < 1 {
    
    port := ":8000"
    fs := http.FileServer(http.Dir(""))
    http.Handle("/", fs)
    fmt.Println("There is no specific port using default. Serving port" + port + " Type your browser http://localhost" + port)
    http.ListenAndServe(port, nil)
    } else {

    port := ":" + preport[0]
    fs := http.FileServer(http.Dir(""))
    http.Handle("/", fs)
    
    fmt.Println("Port Specified Serving port" + port + " Type your browser http://localhost" + port)
    http.ListenAndServe(port, nil)
  }
}