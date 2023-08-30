/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA,lenB:=0,0
	curA,curB:=headA,headB
	var  step int

	for curA!=nil {
		curA = curA.Next
		lenA++
	}
	for curB!=nil {
		curB = curB.Next
		lenB++
	}
	if lenA>lenB{
		step = lenA - lenB
		for i := 0; i < step; i++ {
			headA = headA.Next
		}
	}else {
		step = lenB - lenA
		for i := 0; i < step; i++ {
			headB= headB.Next
		}
	}

	for headB!=headA{
		headA = headA.Next
		headB = headB.Next
	}

	return headA

}

// @lc code=end

