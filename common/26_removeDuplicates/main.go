package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums[:removeDuplicates(nums)])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := len(nums)
	p1, p2 := 1, 1
	for p2 < l {
		if nums[p2] != nums[p2-1] {
			nums[p1] = nums[p2]
			p1++
		}
		p2++
	}
	return p1
}
