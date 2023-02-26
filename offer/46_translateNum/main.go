package main

import "fmt"

func translateNum(num int) int {
	var n []int
	for num != 0 {
		n = append(n, num%10)
		num /= 10
	}
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	m := make(map[int]int)
	m[-1] = 1
	m[0] = 1
	if len(n) < 2 { // 当数组不足两位时，只有一种可能
		return m[0]
	}
	for i := 1; i < len(n); i++ {
		m[i] += m[i-1]
		if n[i-1] != 0 && n[i-1]*10+n[i] < 26 {
			m[i] += m[i-2]
		}
	}
	return m[len(n)-1]
}

func main() {
	fmt.Println(translateNum(12258))
}
