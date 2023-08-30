
/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	for i := 0; i < len(magazine); i++ {
		for j := 0; j < len(ransomNote); j++ {
			if  magazine[i]== ransomNote[j] {
				ransomNote = ransomNote[:j] + ransomNote[j+1:]
				break
			}

		}

	}
	if len(ransomNote) == 0 {
		return true
	}
	return false

}

// @lc code=end

