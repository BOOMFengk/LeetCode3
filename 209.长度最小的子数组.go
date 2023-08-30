/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	i:=0
	n:=len(nums)
	sublength:=n+1
	sum:=0

	for j := 0; j < n; j++ {
		sum+=nums[j]
		for sum>=target{
			sum-=nums[i]
			sublength=min(sublength,j-i+1)
			i++
		}
		
	}
	if sublength == n+1{
		sublength = 0
	}
	return sublength


}
func min(i,j int)int{
	if i<j{
		return i
	}
	return j
}
// @lc code=end

