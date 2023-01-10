package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is  & day is %v\n", day)
	// }

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 2 {
			goto loc
		}

		if roughValue == 4 {
			roughValue++
			continue
		}
		fmt.Println("value is: ", roughValue)
		roughValue++
	}

loc:
	fmt.Println("Hi this is 2 ")
}
