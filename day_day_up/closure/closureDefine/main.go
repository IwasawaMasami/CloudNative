package main

import "fmt"

//闭包

func f2(x, y int) {
	fmt.Println("it's f2")
	fmt.Println(x + y)
}

func dream(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

//定义一个函数对f2进行包装
func main() {
	ret := dream(f2, 100, 200)
	ret()
}
