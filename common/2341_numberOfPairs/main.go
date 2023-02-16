package main

// https://leetcode.cn/problems/maximum-number-of-pairs-in-array/

func numberOfPairs(nums []int) []int {
	m := make(map[int]int)
	answer := []int{0, 0}
	for i := range nums {
		m[nums[i]]++
	}
	for _, v := range m {
		answer[0] += v / 2
		answer[1] += v % 2
	}
	return answer
}

func main() {
	numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2})
}
