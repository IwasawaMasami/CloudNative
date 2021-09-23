package main

import "fmt"

/*
Go语言中 `map`的定义语法如下：
map[KeyType]ValueType
其中，
- KeyType:表示键的类型。
- ValueType:表示键对应的值的类型。
map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
make(map[KeyType]ValueType, [cap])

其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
*/

func main() {
	//var m1 map[string]int
	m1 := make(map[string]int, 10)
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["娜扎"]) //如果不存在这个key就返回对应类型的0值
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("无")
	} else {
		fmt.Println(value)
	}

}
