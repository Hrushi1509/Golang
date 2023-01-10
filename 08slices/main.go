package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"mango", "banana", "peach"}
	fmt.Printf("type is %T\n", fruitList)

	fruitList = append(fruitList, "apple", "tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 555
	highScores[1] = 985
	highScores[2] = 325
	highScores[3] = 400

	highScores = append(highScores, 758, 110)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// delete a particular value

	var courses = []string{"reactjs", "python", "java", "js", "go"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
