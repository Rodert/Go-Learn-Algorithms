/*
 * @Author: JavaPub
 * @Date: 2023-10-28 11:36:49
 * @LastEditors: your name
 * @LastEditTime: 2023-10-28 11:44:59
 * @Description: Here is the JavaPub code base. Search JavaPub on the whole web.
 * @FilePath: \Go-Learn-Algorithms\examples\simple\1\1.go
 */

// 第一道题就用两数求和来开始，两数之和

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, b := range nums {
			if v+b == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{9, 9}
}
