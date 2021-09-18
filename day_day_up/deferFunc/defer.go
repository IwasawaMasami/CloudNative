package main

import "fmt"

func deferDemo() {

	fmt.Println("start")
	defer println("嘿嘿嘿") //defer把它后面的语句延迟执行
	defer println("hhh")
	defer println("xxx") //多个defer按照先进后出的顺序执行
	fmt.Println("end")
}
func main() {
	deferDemo()
}
