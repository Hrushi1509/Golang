package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 & you can open")
	case 2:
		fmt.Println("Dice value is 2 & you can move 2 spaces")
	case 3:
		fmt.Println("Dice value is 3 & you can move 3 spaces")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 & you can move 4 spaces")
	case 5:
		fmt.Println("Dice value is 5 & you can move 5 spaces")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6 & you can play again")
	}
}
