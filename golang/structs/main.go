package main

import "fmt"

func main() {
	fmt.Println("welcom structs")

	devansh := User{"devansh", "devansh@go.dev", true, 18}
	fmt.Println(devansh)
	fmt.Printf("devansh Details are: %+v\n", devansh)
	fmt.Printf("name is %v and email is %v.\n", devansh.Name, devansh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
