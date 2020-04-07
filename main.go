//Nicholas Larsen
//4/7/2020
//asks for the user's name and displays a welcome message using their name
package main

import "fmt"

func nameMessage() {
  var name string
  fmt.Println("Enter your name")
  fmt.Scanln(&name)
  //the user's name
  fmt.Println("Hello",name,"welcome to this program! Thank you for taking the time to run this and leave feedback of any sort!")
  //welcome message
}
func main() {
  nameMessage()
  //recalls the welcome message
}