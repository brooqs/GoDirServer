package main

import (
  "log"
  "net/http"
  "os"
 )

func main() {
  preport := os.Args[1:]

  if len(preport) < 1{
  	log.Println("No port is specified")
  	log.Println("Usage: godirserv port e.g: godirserv 3000 Beware Use high port like 3000 10000 ")
  } else{

  port := ":" + preport[0]
  fs := http.FileServer(http.Dir(""))
  http.Handle("/", fs)
  //fmt.Println(port[0])
  log.Println("Serving port" + port + " Type your browser http://localhost" + port)
  http.ListenAndServe(port, nil)
}
}