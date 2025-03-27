package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup

func inc() {
	i = i + 1
}

func main() {

	fmt.Println(i)
}
