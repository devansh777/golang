package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome my time")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2022, time.May, 22, 23, 22, 22, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
