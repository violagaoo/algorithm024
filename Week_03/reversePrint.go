package main

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var size int = 0
	for p := head; p != nil; p = p.Next {
		size++

	}
	var ret = make([]int, size)
	for ; head != nil; head = head.Next {
		ret[size-1] = head.Val
		size--
	}
	return ret
}

// func main() {

// }
