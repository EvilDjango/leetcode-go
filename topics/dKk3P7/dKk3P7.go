// 剑指 Offer II 032. 有效的变位词
//给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
//
//注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。
//
//
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//示例 2:
//
//输入: s = "rat", t = "car"
//输出: false
//示例 3:
//
//输入: s = "a", t = "a"
//输出: false
//
//
//提示:
//
//1 <= s.length, t.length <= 5 * 104
//s and t 仅包含小写字母
//
//
//进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
//
//
//注意：本题与主站 242 题相似（字母异位词定义不同）：https://leetcode-cn.com/problems/valid-anagram/
//
//通过次数7,442提交次数12,460
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/12/22 2:45 PM
package dKk3P7

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	if s == t {
		return false
	}
	count := map[rune]int{}
	for _, r := range s {
		count[r]++
	}
	for _, r := range t {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}
	return true
}
