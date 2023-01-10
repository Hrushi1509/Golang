package main

import "fmt"

func main() {
	fmt.Println("Struct in Golang")

	// No inheritance in golang; No super or parent

	hw := User{"Hrushi", "hw@go.com", true, 22}
	fmt.Println(hw)

	fmt.Printf("Details are: %+v\n", hw)
	fmt.Printf("Name is %v & Email is %v.", hw.Name, hw.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
