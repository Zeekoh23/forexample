package main

import "fmt"

//DECLARE the variable "y"
//ASSIGN the value 43
// declare and assign = initialization
var y = 43

//DECLARE there is a VARIABLE with the identifier "z"
//and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
//ASSIGNS the ZERO VALUE of TYPE int to "z"
//false for booleans, 0 for integers, 0.0 for floats, "" for strings,
//and nil for pointers, functins, interfaces, slices,
//channels and maps.
var z int

func main() {
	//short declaration operator
	//declare a variable and assign a value
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again: ", y)
}
