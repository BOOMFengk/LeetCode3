/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	n:=len(s)
	if n== 0{
		return false
	}
	j := 0 
	next:= make([]int, len(s))
	next[0] = 0
	for i := 1; i < n; i++ {
		for j>0 && s[i]!=s[j]{
			j = next[j-1]
			
		}
		if s[i]==s[j]{
			j++
		}
		next[i]=j
	}
	// next[n-1] 是最小前后缀的长度
	if next[n-1]!=0&& n%(n-next[n-1])==0{
		return true
	}
	return false
}
// @lc code=end

