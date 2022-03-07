package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("example about files")

	contant := "hi this is Devansh Shah"

	file, err := os.Create("./myFile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, contant)
	if err != nil {
		panic(err)
	}
	fmt.Println("file length is :", length)
	readFile("./myFile.txt")
	defer file.Close()
}

func readFile(fileName string) {
	contant, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contant))
}
