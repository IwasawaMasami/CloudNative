package main

import "fmt"

func main() {

	var s1 []int    //定义存放int类型元素切片
	var s2 []string //定义存放string类型元素切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//初始化
	s1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 = []string{"沙盒", "张江", "坪山区"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	//数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于数组切割，左包含右不包含
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] //[0:4]
	s6 := a1[3:] //[3:len(a1)] 7 9 11 13
	s7 := a1[:]  //[0:len(a1)]
	fmt.Println(s5, s6, s7)
	//切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	//底层数组的从切片第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	//切片再切割
	s8 := s6[3:]
	fmt.Printf("len(s8):%d, cap(s8):%d\n", len(s8), cap(s8))

}
