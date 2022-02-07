package main

import "fmt"

func main() {
	length := 15
	i := 0
	for ; i <= length; i++ {
		fmt.Print(i)
	}
	if length <= 10 {
		fmt.Println("length is less then 10")
	} else {
		fmt.Println("length is more then 10")
	}

	//differ

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
