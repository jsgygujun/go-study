package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // 单个循环条件
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ { // 初始/条件/后续 for 循环
		fmt.Println(j)
	}

	for { // 不带条件的 for 循环将一直重复执行， 直到在循环体内使用了 break 或者 return 跳出循环。
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // 使用 continue 直接进入下一次循环
		}
		fmt.Println(n)
	}
}
