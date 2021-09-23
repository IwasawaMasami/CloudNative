package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

const (
	n1 = 100000000
	n2
	n3
)
const (
	a1 = iota //0
	a2        //1
	a3        //2
)
const (
	b1 = iota
	b2
	_
	b3
)
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	name = "瓜瓜"
	age = 16
	isOk = true
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:\n%s", name)
	fmt.Println(age)

	var s1 string = "滴滴滴滴滴"
	fmt.Println(s1)
	var s2 = "20"
	fmt.Println(s2)
	s3 := "哈哈哈"
	fmt.Println(s3)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
