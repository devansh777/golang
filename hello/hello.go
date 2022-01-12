package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Devansh", "Anand", "Naman", "Anand"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for key, msg := range messages {
		fmt.Println(key + "=>" + msg)
	}

}
