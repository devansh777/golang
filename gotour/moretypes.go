package main

import (
	"fmt"
)

type Mystruct struct {
	name string
	age  int
}

func main() {

	y := Mystruct{"devansh", 25}
	fmt.Println(y.name)

	i := 25
	p := &i
	fmt.Println(*p)
	*p = 50
	fmt.Println(i)
	//array
	var a [2]int
	a[0] = 25
	a[1] = 30
	fmt.Println(a)
	//slices
	prime := [6]int{2, 3, 5, 7, 11, 22}
	fmt.Println(prime[0:6])

}
