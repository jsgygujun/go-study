package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // 声明后却没有给出对应的初始值时，变量将会初始化为 零值 。
	fmt.Println(e)

	f := "short" // 简写
	fmt.Println(f)
}
