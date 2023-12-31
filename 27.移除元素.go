/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	left,right:=0,len(nums)-1
	for left<=right{
		for left<=right&&nums[left]!=val {
			left++
		}
		for  left<=right&&nums[right]==val {
			right--
		}
		if left<=right{
			nums[left]=nums[right]
			left++
			right--
		}
	}
	return left

}
// @lc code=end

