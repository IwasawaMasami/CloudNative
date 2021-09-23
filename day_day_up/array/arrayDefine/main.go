package main

import "fmt"

var a1 [3]bool
var a2 [4]bool

func main() {

	fmt.Printf("%T %T\n", a1, a2)
	fmt.Println(a1, a2)
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	/////////////////////////////////////
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)
	citys := [...]string{"北京", "伤害", "生政"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	fmt.Println(len(citys))
	array()
	arraysum()
}

func array() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}

func arraysum() {

	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)
}
