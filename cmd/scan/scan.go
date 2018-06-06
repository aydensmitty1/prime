package main

import "fmt"
var name string
func main(){
  fmt.Println("What is you name")
  fmt.Scan(&name)
  greet()
}
func greet(){
  fmt.Println("how are you", name)
}
