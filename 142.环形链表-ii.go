package main
/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	//n(a+b)=a+b+c+b
	//假设n为2
	//a=c  当相遇后 头结点和slow节点都往下走 当相遇的时候就是环形链表的入口
	fast,slow := head,head
	for fast!=nil&&fast.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
	
	if fast==slow{
	for head !=slow{
		head=head.Next
		slow=slow.Next
	}
	return head
}
	}
	return nil
}
// @lc code=end

