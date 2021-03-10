package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})              // 创建了一个新的结构体元素
	fmt.Println(Person{name: "Alice", age: 30}) // 初始化一个结构体元素时指定字段名字
	fmt.Println(Person{name: "Fred"})           // 省略的字段将被初始化为零值
	fmt.Println(&Person{name: "Ann", age: 40})  // & 前缀生成一个结构体指针
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 51 // 结构体是可变的
	fmt.Println(sp.age)
}
