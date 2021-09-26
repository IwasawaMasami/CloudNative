package main

import "fmt"

//函数类型
func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

//函数作为参数类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	return a + b
}

//函数作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}
func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	c := f3
	fmt.Printf("%T\n", c)
	f5(f2)
	f7 := f5(f2)
	fmt.Printf("%T\n", f7)

}
