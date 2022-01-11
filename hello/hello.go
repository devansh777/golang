package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Devansh")
	fmt.Println(message)
}
