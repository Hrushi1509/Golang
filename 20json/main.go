package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodedJson()
	DecoderJson()
}

func EncodedJson() {
	ytCourses := []course{
		{"ReactJS Bootcamp", 299, "youtube.com", "abc@123", []string{"web-dev", "frontend"}},
		{"MERN Bootcamp", 499, "youtube.com", "bcd@123", []string{"Full-Stack", "web-dev"}},
		{"NodeJS Bootcamp", 199, "youtube.com", "xyz@123", nil},
	}

	finalJSON, err := json.MarshalIndent(ytCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)

}

func DecoderJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"Platform": "youtube.com",
			"tags": ["web-dev","frontend"]
		}
	`)

	var ytCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &ytCourse)
		fmt.Printf("%#v\n", ytCourse)
	} else {
		fmt.Println("Json is not valid")
	}

	// Some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}
