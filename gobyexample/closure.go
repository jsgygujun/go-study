package main

import "fmt"

// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数。这个返回的函数使用闭包的方式隐藏变量 i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())
	fmt.Println(newNextInt())
}
