package main
/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for k1, _:= range nums {
		for k2 := k1+1; k2 < len(nums); k2++ {
			if nums[k1]+nums[k2]==target{
			return []int{k1,k2}
			}	
		}
	}
	return []int{}
}
// @lc code=end

