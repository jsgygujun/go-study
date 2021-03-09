package main

import "fmt"

func main() {
	var a [5]int // 存放 5 个 int 元素的数组 a
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println(len(a)) // 内置函数 len 可以返回数组的长度

	b := [5]int{1, 2, 3, 4, 5} // 在一行内声明并初始化一个数组
	fmt.Println(b)

	var twoD [2][3]int // 多维的数据结构。
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
