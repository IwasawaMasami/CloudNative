package main

import "fmt"

func main() {
	var s string = "我爱你，小可爱"
	for i, j := range s {
		fmt.Printf("%d %c\n", i, j)
	}

}