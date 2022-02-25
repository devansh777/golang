package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to use Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our services:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thnks for rating, ", input)
}
