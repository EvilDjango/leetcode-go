// 剑指 Offer II 019. 最多删除一个字符得到回文
//给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
//
//
//
//示例 1:
//
//输入: s = "aba"
//输出: true
//示例 2:
//
//输入: s = "abca"
//输出: true
//解释: 可以删除 "c" 字符 或者 "b" 字符
//示例 3:
//
//输入: s = "abc"
//输出: false
//
//
//提示:
//
//1 <= s.length <= 105
//s 由小写英文字母组成
//
//
//注意：本题与主站 680 题相同： https://leetcode-cn.com/problems/valid-palindrome-ii/
//
//通过次数13,526提交次数29,511
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/21 下午3:35
package RQku0D

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left:right]) || isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r && s[l] == s[r] {
		l++
		r--
	}
	return l >= r
}
