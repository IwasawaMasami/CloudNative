package main

import "fmt"

func main() {
	m1 := make(map[string]int, 10)
	m1["理想"] = 18
	m1["jiwuming"] = 35

	for k := range m1 {
		fmt.Println(k)
	}
	for _,v := range m1 {
		fmt.Println(v)
	}
	delete(m1,"jiwuming")
	fmt.Println(m1)

}