package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointer")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of ptr ", *ptr)
	fmt.Println("value of ptr ", ptr)

	*ptr = *ptr * 2
	fmt.Println(myNumber)
}
