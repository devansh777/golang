package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println("My favorite number is", rand.Intn(20))
	fmt.Println(add(20, 25))
	fmt.Println(math.Pi)
	a, b := swap("anand", "Devansh")
	fmt.Println(a)
	fmt.Println(b)
	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
}

func add(a int, b int) int {
	return a + b
}
func swap(p, q string) (string, string) {
	return q, p
}
