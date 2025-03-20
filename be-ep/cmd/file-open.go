package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Open("go.mod")
	if err != nil {
		panic(err) // throw Golang error
	}

	//defer file.Close()

	file2, err := os.Open("file-open.go")
	if err != nil {
		panic(err)
	}
	//defer file2.Close()

	defer func() {
		file.Close()
		file2.Close()
	}()
}
