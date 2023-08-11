package main

import (
  "fmt"
)

func main() {
  mat := make([][]int, 0)
  mat = append(mat, []int{1,2,3})
  mat = append(mat, []int{4,5,6})
  mat = append(mat, []int{7,8,9})
  fmt.Println("sum is %v", diagonalSum(mat))
}

func diagonalSum(mat [][]int) int {
  l := len(mat)
  var sum int
  for i := 0; i < l; i++ {
    sum += mat[i][i] + mat[i][l-i-1]
  }
  if l % 2 == 1 {
    sum -= mat[l/2][l/2]
  }
  return sum
}
