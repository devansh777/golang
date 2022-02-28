package main

import "fmt"

func main() {
	fmt.Println("welcom to array")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "tometo"
	fruitList[3] = "peach"

	fmt.Println("fruit list is: ", fruitList)
	fmt.Println("fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("veg List is: ", len(vegList))

}
