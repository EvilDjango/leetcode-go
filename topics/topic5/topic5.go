// 5. 最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"
//示例 3：
//
//输入：s = "a"
//输出："a"
//示例 4：
//
//输入：s = "ac"
//输出："a"
//
//
//提示：
//
//1 <= s.length <= 1000
//s 仅由数字和英文字母（大写和/或小写）组成
//通过次数726,303提交次数2,051,688
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/16/21 2:01 PM
package topic5

import "strings"

// Manacher算法
func longestPalindrome(s string) string {
	sb := strings.Builder{}
	sb.WriteByte('#')
	for _, r := range s {
		sb.WriteRune(r)
		sb.WriteByte('#')
	}
	s = sb.String()

	center, right := 0, 0
	start, end := 0, 2
	lens := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		length := 0
		if i < right {
			length = min(lens[2*center-i], right-i)
		}
		l, r := i-1-length, i+1+length
		for l >= 0 && r < len(s) && s[l] == s[r] {
			length++
			l--
			r++
		}
		lens[i] = length
		if length*2+1 > end-start+1 {
			start, end = i-length, i+length
		}
		if i+length > right {
			center = i
			right = i + length
		}
	}
	sb = strings.Builder{}
	for i := start + 1; i < end; i += 2 {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
