package main

// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/submissions/

var m = make(map[int]int)

func fib(n int) int {
	m[0] = 0
	m[1] = 1
	for i := 2; i <= n; i++ {
		m[i] = (m[i-1] + m[i-2]) % 1000000007
	}
	return m[n]
}
