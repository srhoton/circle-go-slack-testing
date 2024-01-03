package main

import (
  "fmt"
  "os"
  "math/rand"
)

func main() {
  fmt.Println("hello world")
  xc := rand.Intn(100)
  if xc%2 == 0 { 
    xc = 100
  } else {
    xc = 0
  }
  fmt.Println("exiting with code: ", xc)
  os.Exit(xc)
}
