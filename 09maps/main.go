package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println(languages)
	fmt.Println("py short for ", languages["py"])

	delete(languages, "py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
