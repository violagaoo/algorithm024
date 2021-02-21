package main

//剑指 Offer 68 - II. 二叉树的最近公共祖先

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	lret := lowestCommonAncestor(root.Left, p, q)
	rret := lowestCommonAncestor(root.Right, p, q)
	//当 left 为空 ，right 不为空 ：p,q 都不在 rootroot 的左子树中，直接返回 rightright 。具体可分为两种情况：
	//p,qp,q 其中一个在 rootroot 的 右子树 中，此时 rightright 指向 p（假设为 p）；
	//p,qp,q 两节点都在 rootroot 的 右子树 中，此时的 rightright 指向 最近公共祖先节点 ；
	//当 leftleft 不为空 ， rightright 为空 ：与情况 3. 同理；

	if lret != nil && rret != nil {
		return root
	} else if lret != nil {
		return lret
	} else if rret != nil {
		return rret
	}
	return nil
}

// func mian7() {

// }
