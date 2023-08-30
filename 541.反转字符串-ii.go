
/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	ss :=[]byte(s)
	for i := 0; i < len(ss); i+=2*k {
		if i+k < len(ss){
			reverseString(ss[i:i+k])
		}else{
			reverseString(ss[i:])
		}
	}
	return string(ss)
}

func reverseString(s []byte)  {
	left,right:=0,len(s)-1
	for left < right {
		s[left],s[right] = s[right],s[left]
		left++
		right--
	}

}

// @lc code=end

