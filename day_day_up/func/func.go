package main

import "fmt"

//函数 是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中给它起一个名字，每次用到的时候就可以直接用函数名字

func sum(x int, y int) (ret int) {
	return x + y
}

/*
//没有返回值的函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int {
	return 3
}

//返回值可以命名也可以不命名
//命名返回值就相当于函数中声明的一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以return后省略
}

//参数的类型简写
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

//可变长参数
//必须放在函数参数最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
*/
//Go语言中没有默认参数这个概念
func main() {

	r := sum(1, 2)
	fmt.Println(r)

}
