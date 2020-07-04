package main

import "fmt"

var a int

type hotdog int//creating my own data type

var b hotdog//declaring a variable with my datatype

func main() {
	a = 42s
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n ", a)//format option
	fmt.Println(b)
	fmt.Printf("%T\n ", b)

	//a = b
	//fmt.Println(a)
	//fmt.Printf("%T\n ", a)

}
