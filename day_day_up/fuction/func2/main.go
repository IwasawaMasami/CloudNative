package main

import "fmt"

func f1() {
	fmt.Println("沙盒")
}
func f2(name string) {
	fmt.Println("hello",name)
}

func f3(x int,y int)int {
	sum := x + y
	return sum
}
func main() {
	f1()
	f2("理想")
	ret := f3(100,200)
	fmt.Println(ret)
	
}