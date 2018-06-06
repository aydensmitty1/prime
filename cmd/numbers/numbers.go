package main

import ("github.houston.softwaregrp.net/CSB/gotest/pkg/prime"
 "fmt"
)
var max int
func main(){
  fmt.Println("Pick any number")
  fmt.Scan(&max)
  prime.Pnums(max)

}
