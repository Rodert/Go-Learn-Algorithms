package main

import "fmt"

func main() {
	fmt.Println(twoSumV2([]int{2, 7, 11, 15}, 9))
}

func twoSumV2(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
