package main

import "fmt"

func main() {

	s1 := []string{"北京", "伤害", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州" //数组越界问题
	fmt.Println(s1)
	//使用append必须用原来的切片接收返回值
	s1 = append(s1, "广州")
	fmt.Println(s1)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"GG", "MM", "HH"}
	s1 = append(s1, ss...)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
