package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { // 数组的range返回的第一个是元素的下标，第二个是元素的值
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a:": "apple", "b": "banana"} // range 在 map 中迭代键值对
	for k, v := range kvs {
		fmt.Printf("%s - %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
