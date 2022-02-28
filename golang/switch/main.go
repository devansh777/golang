package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("dice game with switch Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	default:
		fmt.Println("not possible")
	}
}
