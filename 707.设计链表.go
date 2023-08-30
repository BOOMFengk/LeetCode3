package main

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
type Node struct{
	Val int
	Next *Node
}	


type MyLinkedList struct {
	dummy *Node
	size int
}	



func Constructor() MyLinkedList {
	return MyLinkedList{&Node{},-1}
	
}

func (this *MyLinkedList) Get(index int) int {
	if index<0||index>this.size{
		return -1
	}
	cur:=this.dummy
	for i := 0; i <= index; i++ {
		cur = cur.Next
		
	}
	return cur.Val
	
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newNode:=&Node{val,this.dummy.Next}
	this.dummy.Next = newNode
	this.size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
	newNode:=&Node{val,nil}
	cur:= this.dummy
	this.size++
	for cur.Next!=nil {
		cur=cur.Next
		
	}
	cur.Next = newNode
	
	
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index<=0{
		this.AddAtHead(val)

	}else if index==this.size+1 {
		this.AddAtTail(val)

		
	}else if index>this.size+1 {
		return
		
	}else {
		cur:=this.dummy
		this.size++
		for i := 0; i < index; i++ {
			cur = cur.Next
		}
		newNode:=&Node{val,cur.Next}
		cur.Next =newNode
	}
}




func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index<0 || index >this.size{
		return
	}
	this.size--
	var prev *Node
	cur:=this.dummy
	for i := 0; i <= index; i++ {
		prev = cur
		cur = cur.Next
	}
	prev.Next = cur.Next
}



/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

