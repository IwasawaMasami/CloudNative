package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	//1.判断字符串中的汉字数量
	//难点判断一个字符是否为汉字

	s1 := "Hello沙盒" //单词出现次数
	s2 := "how do you do"
	//依次拿到字符串中的字符
	var count int
	for _, c := range s1 {
		//判断当前字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	s3 := strings.Split(s2, " ")   //字符串切割按照空格得到切片
	m1 := make(map[string]int, 10) //遍历切片存储到一个map
	for _, w := range s3 {         //1. 如果原来map中不存在这个key 2.如果map中存在这个key，那么出现次数+1
		fmt.Println(w)
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		for key, value := range m1 {
			fmt.Println(key, value)
		}
	}

	fmt.Println(count)

}
