package main

import "fmt"

func main() {
	fmt.Println("function practice")

	result, msg := adder(1, 2, 3, 4, 5, 5)
	fmt.Printf("%v : %v", msg, result)

}

func adder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "sum of all values"
}
