// 剑指 Offer II 018. 有效的回文
//给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
//
//本题中，将空字符串定义为有效的 回文串 。
//
//
//
//示例 1:
//
//输入: s = "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//示例 2:
//
//输入: s = "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//提示：
//
//1 <= s.length <= 2 * 105
//字符串 s 由 ASCII 字符组成
//
//
//注意：本题与主站 125 题相同： https://leetcode-cn.com/problems/valid-palindrome/
//
//通过次数15,475提交次数29,906
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/22 上午9:07
package XltzEq

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isDigitOrLetter(s[l]) {
			l++
		}
		for r > l && !isDigitOrLetter(s[r]) {
			r--
		}
		if !equalsIgnoreCase(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isDigitOrLetter(b byte) bool {
	// 数字
	if b >= 48 && b <= 57 {
		return true
	}
	// 大写字母
	if b >= 65 && b <= 90 {
		return true
	}
	// 小写字母
	if b >= 97 && b <= 122 {
		return true
	}
	return false
}

func equalsIgnoreCase(r1, r2 byte) bool {
	if r1 == r2 {
		return true
	}
	if r1 > r2 {
		r1, r2 = r2, r1
	}

	return r1 >= 'A' && r1 <= 'Z' && r2-r1 == 32
}
