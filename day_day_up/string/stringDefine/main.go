package main

import "fmt"

func main() {
	s3 := "落"
	s4 := "D:\\Go\\src"
	s5 := "我爱你"
	fmt.Printf("%s\n", s3)
	fmt.Printf("%s\n", s4)
	fmt.Print(len(s3))
	fmt.Println()
	fmt.Print(len(s4))
	fmt.Println()
	fmt.Print(len(s5))
	fmt.Println()
	fmt.Printf("%s\n", s5)
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	changeString()
}
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
