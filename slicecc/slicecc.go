package main

import(
	"fmt"
)

func main(){
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Println("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	
}