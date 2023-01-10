package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10.5
	var result string

	if loginCount < 10 {
		result = "less than 10"
	} else if loginCount > 10 {
		result = "more than 10"
	} else {
		result = "exactly 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("more than 10")
	}
}
