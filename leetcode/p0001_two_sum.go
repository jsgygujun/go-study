package main

import "fmt"

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, val := range nums {
		matchedVal := target - val
		matchedIdx, ok := lookup[matchedVal]
		if ok == true {
			return []int{matchedIdx, i}
		}
		lookup[val] = i
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
