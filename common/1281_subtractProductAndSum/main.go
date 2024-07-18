package main

// https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
import (
	"fmt"
	
)

func main() {
	num := 234
	fmt.Println(subtractProductAndSum(num))
}

func subtractProductAndSum(n int) int {
	list := make([]int, 0)
	for n > 0 {
		list = append(list, n%10)
		n /= 10
	}
	product, sum := 1, 0
	for i := range list {
		product *= list[i]
		sum += list[i]
	}
	return product - sum
}
