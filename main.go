package main

import "fmt"

func nameMessage() {
  var name string
  fmt.Println("Enter your name")
  fmt.Scanln(&name)
  fmt.Println("Hello",name,"welcome to this program! Thank you for taking the time to run this and leave feedback of any sort!")
}
func main() {
  nameMessage()
}