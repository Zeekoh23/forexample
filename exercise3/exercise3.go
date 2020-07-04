package main

import "fmt"

type isaac int
var x isaac

func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x := 42
	fmt.Println(x)


}