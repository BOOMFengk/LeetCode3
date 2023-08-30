/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	record:= [26]int{}
	for _, v := range s {
		record[v-rune('a')]++
		
	}
	for _, v := range t {
		record[v-rune('a')]--
		
	}
	return record == [26]int{}

}


// @lc code=end

