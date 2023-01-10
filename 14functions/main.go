package main

import "fmt"

func main() {
	fmt.Println("functions")

	res := adder(5, 8)
	fmt.Println(res)

	res2, message := adder2(4, 8, 7, 6, 4)
	fmt.Println(res2)
	fmt.Println(message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func adder2(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "this is sum"
}
