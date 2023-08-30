package main
/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	mp1:=make(map[int]struct{},0)
	result := make([]int, 0)
	for _, v := range nums1 {
		if _,ok:=mp1[v];!ok{//去重如果有多个这个只取一个
			mp1[v]=struct{}{}
		}
		
	}
	for _, v := range nums2 {
		if _,ok:=mp1[v];ok{//去重
			result=append(result,v)	
			delete(mp1,v)		
		}
	}
	return result
	
}
// @lc code=end

