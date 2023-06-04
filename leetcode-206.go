package main

// 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	// 递归解法
	// 很好理解，就是两两节点反转即可
	//return reverse(nil, head)

	// 双指针
	// 其实主要就是需要下一个节点就可以了
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverse(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	next := head.Next
	next.Next = pre
	return reverse(head, next)
}
