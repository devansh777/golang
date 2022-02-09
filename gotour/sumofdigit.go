package main

import "fmt"

func main() {
	sum := sumDigi(555)
	fmt.Println("sum of enter digit is:", sum)
}

func sumDigi(no int) int {
	total := no
	for total >= 10 {
		lastDigi := total % 10
		firstDigi := total / 10
		total = lastDigi + firstDigi
	}
	return total
}
