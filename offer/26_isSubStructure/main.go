package main

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=j7ivt3c

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isCheck bool

var B *TreeNode

func isSubStructure(a, b *TreeNode) bool {
	if b == nil {
		return false
	}
	B = b
	return check(a, b)
}

func check(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		if check(a.Left, B) || check(a.Right, B) {
			isCheck = true
		}
	} else {
		if check(a.Left, b.Left) && check(a.Right, b.Right) {
			return true
		}
	}
	return isCheck
}

func main() {
	a := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	b := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
	}
	isSubStructure(a, b)
}
