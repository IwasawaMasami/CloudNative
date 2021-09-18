package main

import "fmt"

func main() {

	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n",c1,c2)
	fmt.Printf("c2:%d\n",c2)
	fmt.Println(len(s2))
}