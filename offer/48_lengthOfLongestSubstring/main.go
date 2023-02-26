package main

import "fmt"

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=j7ivt3c

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	m := make(map[string]int)
	var l, r int
	for i, c := range s {
		if _, ok := m[string(c)]; !ok {
			m[string(c)] = i

		} else {
			if r-l > max {
				max = r - l
			}
			if l < m[string(c)]+1 {
				l = m[string(c)] + 1
			}
			m[string(c)] = i
		}
		r++
	}
	if r-l > max {
		max = r - l
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("sdqweuqioweioionsjkdfilsduhfgl"))
}
