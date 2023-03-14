package main

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=j7ivt3c

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int
var tar int

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	tar = target
	dfs(root, 0, []int{})
	return res
}

func dfs(node *TreeNode, sum int, l []int) {
	if node == nil || sum+node.Val > tar {
		return
	}
	l = append(l, node.Val)
	if node.Left != nil {
		dfs(node.Left, sum+node.Val, l)
	}
	if node.Right != nil {
		dfs(node.Right, sum+node.Val, l)
	}
	if sum == tar && node.Left == nil && node.Right == nil {
		res = append(res, l)
	}
}
