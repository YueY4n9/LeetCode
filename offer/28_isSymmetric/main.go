package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var rootList = &[]interface{}{}
	var mirrorList = &[]interface{}{}
	dfs(root, rootList)
	mirror(root)
	dfs(root, mirrorList)
	if len(*rootList) != len(*mirrorList) {
		return false
	}
	for i := range *rootList {
		if (*rootList)[i] != (*mirrorList)[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, rootList *[]interface{}) {
	if root == nil {
		*rootList = append(*rootList, nil)
		return
	}
	*rootList = append(*rootList, root.Val)
	dfs(root.Left, rootList)
	dfs(root.Right, rootList)
}

func mirror(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	mirror(root.Left)
	mirror(root.Right)
}

func main() {
	a := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	isSymmetric(a)
}
