package main

import "fmt"

func main() {
	// & 取地址
	n := 18
	p := &n
	fmt.Println(n)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	// * 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Println("*********************************************************")
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}