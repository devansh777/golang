package main

import "fmt"

func main() {
	devansh := Users{"devansh", 18, true}
	fmt.Println(devansh.Name)
	devansh.getAge()
	devansh.getStatus(devansh.status)
}

type Users struct {
	Name   string
	Age    int
	status bool
}

func (u Users) getAge() {
	u.Age = 18
	fmt.Println(u.Age)
}

func (u Users) getStatus(status bool) {
	u.status = status
	fmt.Println(u.status)
}
