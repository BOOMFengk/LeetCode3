/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	m:=make(map[int]bool, 0)
	for n!=1&&!m[n] {// !m[n]判断不循环 如果出现相同的数就代表循环了
		n,m[n]=getSum(n),true
		
	}
	return n == 1
}

func getSum(i int)int{
	sum:=0
	for i>0{
		sum += (i%10)*(i%10)
		i/=10
	}
	return sum

}
// @lc code=end

