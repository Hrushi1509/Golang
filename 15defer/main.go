package main

import "fmt"

func main() {
	fmt.Println("defer in golang")

	defer fmt.Println("hello")
	fmt.Println("world")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
