/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy:=&ListNode{Next:head}
	cur:=dummy
	for cur!=nil&&cur.Next!=nil{
		if cur.Next.Val==val{
			cur.Next=cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
// @lc code=end

