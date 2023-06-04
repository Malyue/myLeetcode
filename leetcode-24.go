package main

// 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
// 示例：
// 给定 1->2->3->4，你应该返回 2->1->4->3

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 先试一下递归
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//next := head.Next
	//// 第一个的下一个会变成第三个,每次传入的都是两个中的第一个或者为空
	//head.Next = swapPairs(next.Next)
	//next.Next = head
	//return next

	// 非递归解法，双指针
	// 声明一个虚拟头节点
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		// 保存下一组节点的起始位置
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		pre = head
		head = next
	}
	return dummy.Next
}
