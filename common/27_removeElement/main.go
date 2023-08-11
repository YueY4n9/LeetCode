package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(nums[:removeElement(nums, val)])
}

func removeElement(nums []int, val int) int {
	p1, p2 := 0, 0
	l := len(nums)
	for p2 < l {
		if nums[p2] == val {
			p2++
			continue
		} else {
			nums[p1] = nums[p2]
			p1++
			p2++
		}
	}
	return p1
}
