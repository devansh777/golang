package main

import "fmt"

func main() {
	fmt.Println("welcome maps")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("list of all lang: ", languages)
	fmt.Println("js for : ", languages["js"])

	delete(languages, "rb")
	fmt.Println("list of all lang: ", languages)

	// loops are in map

	for key, values := range languages {
		fmt.Printf("for key %v, value is %v \n", key, values)
	}

}
