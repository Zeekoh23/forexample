package main

import "fmt"

var y string
var z int

func main() {
	//DECLARE a variable to be of a certain TYPE
	//and then ASSIGN a value of that type to the variable

	fmt.Println("printing the value of y ", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T ", y)
	fmt.Println(z)
	fmt.Printf("%T", z)

}
