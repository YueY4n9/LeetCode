package main

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	for l, r := 0, len(nums)-1; l < r; {
		if nums[l]%2 == 1 {
			l++
			continue
		}
		if nums[r]%2 == 0 {
			r--
			continue
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
