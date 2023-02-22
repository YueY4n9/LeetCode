package main

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=j7ivt3c

var m = make(map[int]int)

func numWays(n int) int {
	m[0] = 1
	m[1] = 1
	for i := 2; i <= n; i++ {
		m[i] = (m[i-1] + m[i-2]) % 1000000007
	}
	return m[n]
}
