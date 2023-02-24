package main

import "fmt"

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=j7ivt3c

func maxValue(grid [][]int) int {
	var l = make([][]int, len(grid))
	for i := range l {
		l[i] = make([]int, len(grid[0]))
	}
	copy(l, grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				l[i][j] = l[i][j] + l[i][j-1]
				continue
			}
			if j == 0 {
				l[i][j] = l[i][j] + l[i-1][j]
				continue
			}
			l[i][j] = max(l[i][j]+l[i-1][j], l[i][j]+l[i][j-1])
		}
	}
	return l[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	test := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(maxValue(test))
}
