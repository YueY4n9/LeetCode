package main

import(
  "fmt"
)

func main() {
  nums1 := []int{1,2,3,0,0,0}
  nums2 := []int{2,5,6}
  m, n := 3, 3
  merge(nums1, m, nums2, n)
  fmt.Println(nums1, nums2)
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
  p1, p2 := 0, 0
  res := make([]int, 0, m+n)
  for {
    if p1 == m {
      res = append(res, nums2[p2:]...)
      break
    }
    if p2 == n {
      res = append(res, nums1[p1:]...)
      break
    }
    if nums1[p1] > nums2[p2] {
      res = append(res, nums2[p2])
      p2++
    }else {
      res = append(res, nums1[p1])
      p1++
    }
  }
  copy(nums1, res)
}
