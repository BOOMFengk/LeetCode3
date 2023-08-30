/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	i,j,k:=0,len(nums)-1,len(nums)-1
	result:=make([]int, len(nums))
	for i<=j{
		a:=nums[i]*nums[i]
		b:=nums[j]*nums[j]
		if a>b{
			result[k]=a
			i++
		}else {//包含==的情况
			result[k]=b
			j--
		}
		k--

	}
	return result

}
// @lc code=end

