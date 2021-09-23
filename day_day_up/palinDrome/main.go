package main

import "fmt"

func main() {
	//回文判断
	//上海自来水来自海上 s[0] s[len(s)-1]
	//山西运煤车煤运西山
	//黄山落叶松叶落山黄
	ss := "山西运煤车煤运西山"
	r := make([]rune, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		// 山 ss[0] s[len(s)-1]
		// 西 ss[1] s[len(s)-1-1]
		// 运 ss[2] s[len(s)-1-2]
		// 煤 ss[3] s[len(s)-1-3]
		//...
		//c   ss[i] s[len(s)-1-i]
		if r[i] != r[len(ss)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
