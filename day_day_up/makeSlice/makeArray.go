package main

import "fmt"

func main() {

	s1 := make([]int, 5, 10)
	fmt.Printf("%d", len(s1))
	fmt.Println(s1)
}