package main

import "fmt"

func main() {
	// & 取地址
	n := 18
	p := &n
	fmt.Println(n)
	fmt.Println(p)
	fmt.Printf("%T\n",p)
	// * 根据地址取值
	m := *p
	fmt.Println(m)
}