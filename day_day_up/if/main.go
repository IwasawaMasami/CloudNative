package main

import "fmt"

func main() {
	age := 15
	if age > 18 {
		fmt.Printf("yes")
	} else if age < 16{
		fmt.Printf("no")
	} else{
		fmt.Printf("%d", age)
	}
	if age :=30; age > 18{
		fmt.Println("23333")
	}else {
		fmt.Println("44444")
	}
}