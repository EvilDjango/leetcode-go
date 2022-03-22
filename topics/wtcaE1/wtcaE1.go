// 剑指 Offer II 016. 不含重复字符的最长子字符串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成
//
//
//注意：本题与主站 3 题相同： https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//
//通过次数16,357提交次数34,219
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/22 上午10:25
package wtcaE1

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	maxLen := 0
	window := map[byte]bool{}
	for r < len(s) {
		b := s[r]
		for window[b] {
			delete(window, s[l])
			l++
		}
		window[b] = true
		if r-l > maxLen {
			maxLen = r - l
		}
		r++
	}
	return maxLen
}
