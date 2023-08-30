/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp1 := make(map[int]int)
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mp1[v1+v2]++
		}
	}
	//v1+v2+v3+v4=0    v1+v2=-v3-v4
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += mp1[-v3-v4]

		}
	}
	return count

}

// @lc code=end

