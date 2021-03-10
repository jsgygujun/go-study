package main

import "fmt"

// 接收两个 int 变量，返回一个 int 结果
func plus(a int, b int) int {
	return a + b
}

// 多个返回值
func values() (int, int) {
	return 3, 7
}

// 变参函数
func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	a, b := values()
	fmt.Println(a, b)

	_, c := values() // 忽略第一个返回值
	fmt.Println(c)

	total := sum(1, 2, 3, 4)
	fmt.Println("sum:", total)

	nums := []int{1, 2, 3, 4, 5}
	total = sum(nums...) // slice 作为变参使用
	fmt.Println("sum:", total)
}
