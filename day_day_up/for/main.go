package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	forDemo3()
	forDemo2()
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}
