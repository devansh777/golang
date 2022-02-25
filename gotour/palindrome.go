package main

import "fmt"

func main() {
	fmt.Println(isPelendrome(10))
}

func isPelendrome(no int) bool {
	orignalNo := no
	pelno := 0
	for no >= 10 {
		temp := no % 10
		no = no / 10
		pelno = (pelno * 10) + temp
	}
	pelno = (pelno * 10) + no
	if pelno == orignalNo {
		return true
	} else {
		return false
	}
}
