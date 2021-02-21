package main

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var last *TreeNode

func isValidBST(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root
	return dfs(root.Right)
}

// func mian5() {

// }
