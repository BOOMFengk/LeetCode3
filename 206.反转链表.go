package main

import "github.com/go-logr/logr/funcr"

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
     return	 help(nil,head)

}
func help(prev,head *ListNode)*ListNode{
	if head==nil{
		return prev
	}

		next:=head.Next
		head.Next = prev

	return help(head,next)
}

// @lc code=end

