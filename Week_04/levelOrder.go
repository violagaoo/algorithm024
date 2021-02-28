package mian

//leetcode 102. 二叉树的层序遍历
/**
 * Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret = make([][]int, 0)
	dfs(root, 0)
	return ret
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(ret) {
		ret = append(ret, []int{})
	}
	ret[level] = append(ret[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
