package main

import "fmt"

// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=j7ivt3c

func movingCount(m int, n int, k int) int {
	var ll [][]bool
	var res int
	ll = make([][]bool, m)
	for i := 0; i < m; i++ {
		ll[i] = make([]bool, n)
	}
	// 1. 计算符合条件的格子坐标
	for i := range ll {
		for j := range ll[i] {
			if i == 0 && j == 0 {
				ll[i][j] = true
				res++
				continue
			}
			if i == 0 {
				ll[i][j] = ll[i][j-1]
			} else if j == 0 {
				ll[i][j] = ll[i-1][j]
			} else {
				ll[i][j] = ll[i][j-1] || ll[i-1][j]
			}
			ll[i][j] = ll[i][j] && check(i, j, k)
			if ll[i][j] {
				res++
			}
		}
	}
	return res
}

func check(i, j, k int) bool {
	return countNumBit(i)+countNumBit(j) <= k
}

func countNumBit(n int) int {
	var res int
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}

func main() {
	fmt.Println(movingCount(3, 1, 0))
}
