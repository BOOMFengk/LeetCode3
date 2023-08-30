package main



/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy:=&ListNode{Next: head}
	first:=dummy
	for head!=nil&&head.Next!=nil{
	second :=head
	third:=head.Next
	next:=head.Next.Next
	first.Next=third
	third.Next=second
	second.Next = next
	first = second
	head=next
	}
	return dummy.Next
}
// @lc code=end

