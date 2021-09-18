package main

import "fmt"

func main() {

	doubleArray()

}
func doubleArray() {
	var num = []int{1, 3, 5, 7, 8}
	sum := 8
	for i := 0; i < len(num); i++ { //len长度为5
		for j := i; j < len(num); j++ {
			fmt.Printf("J的索引为:%d I的索引为:%d\n", j, i)
			if num[i]+num[j] == sum {
				fmt.Println("****************************************************")
				fmt.Printf("成功找到和等于8的两位索引: J的索引为:%d I的索引为:%d *\n", j, i)
				fmt.Println("****************************************************")
			}
		}
	}
}
