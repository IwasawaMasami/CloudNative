package main

import "fmt"

func main() {
	// * 根据地址取值
    var a = new(int) //空指针
	fmt.Println(a)
	fmt.Println(*a)
	
}