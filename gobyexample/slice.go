package main

import "fmt"

func main() {
	s := make([]string, 3) // 创建了一个长度为 3 的 string 类型的 slice（初始值为零值）
	fmt.Println("emp:", s) // emp: [  ]

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d") // 内建函数 append， 该函数会返回一个包含了一个或者多个新值的 slice
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // len 返回 slice 的长度
	copy(c, s)                  // 内建 copy 函数
	fmt.Println("cpy:", c)

	l := s[2:5] // 支持通过 slice[low:high] 语法进行“切片”操作，包含low，不包含high
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // 声明并初始化 slice
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
