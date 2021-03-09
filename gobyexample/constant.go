package main

import (
	"fmt"
	"math"
)

const s string = "constant string" // 包级别常量

func main() {
	fmt.Println(s)

	const n = 50000000 // 函数级别常量
	const d = 2e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
