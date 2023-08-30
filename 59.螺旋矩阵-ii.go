/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	ans:=make([][]int, n)
	for v := range ans {
		ans[v]=make([]int, n)
		
	}
	up,down,left,right:=0,n-1,0,n-1
	cur:=1
	for cur<=n*n{
		for i := left; i <= right; i++ {
			ans[up][i]=cur
			cur++
		}
		up++

		for i := up; i <= down; i++ {
			ans[i][right]=cur
			cur++
		}
		right--
		for i := right; i >= left; i-- {
			ans[down][i]=cur
			cur++
		}
		down--
		for i := down; i >= up; i-- {
			ans[i][left]=cur
			cur++
		}
		left++
	}
	return ans


}
// @lc code=end

