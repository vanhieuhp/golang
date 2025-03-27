package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f3,", err)
		}
	}()
	f1()
}

func f1() {
	fmt.Println("F1 runs oke")
	f2()
}

func f2() {
	fmt.Println("F2 runs oke")

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("Recovered in f3,", err)
	//	}
	//}()

	f3()
}

func f3() {
	fmt.Println("F3 runs oke")
	panic("error from f3")

	fmt.Println("OMG! I can't print")
}
