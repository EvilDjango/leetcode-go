// 剑指 Offer II 020. 回文子字符串的个数
//给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
//示例 1：
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//示例 2：
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//提示：
//
//1 <= s.length <= 1000
//s 由小写英文字母组成
//
//
//注意：本题与主站 647 题相同：https://leetcode-cn.com/problems/palindromic-substrings/
//
//通过次数13,024提交次数18,066
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/21 下午1:35
package a7VOhD

import "strings"

// 动态规划 依次枚举长度从1到n的回文串
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	count := 0
	for length := 2; length <= n; length++ {
		for i := 0; i < n; i++ {
			end := i + length - 1
			dp[i][end] = s[i] == s[end] && (length == 2 || dp[i+1][end-1])
			if dp[i][end] {
				count++
			}
		}
	}
	return count
}

// 中心扩展法 从2n-1个位置分别扩展
func countSubstrings2(s string) int {
	n := len(s)
	count := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			l--
			r++
		}
	}
	return count
}

// Manacher算法
func countSubstrings3(s string) int {
	sb := strings.Builder{}
	sb.WriteRune('^')
	for _, r := range s {
		sb.WriteRune('#')
		sb.WriteRune(r)
	}
	sb.WriteRune('#')
	sb.WriteRune('$')
	s = sb.String()
	n := len(s)
	lens := make([]int, n)
	pivot, right := -1, -1
	count := 0
	for i := 2; i < n-2; i++ {
		if i <= right {
			lens[i] = min(lens[2*pivot-i], right-i+1)
		}
		for s[i-lens[i]] == s[i+lens[i]] {
			lens[i]++
		}
		if i+lens[i]-1 > right {
			pivot, right = i, i+lens[i]-1
		}
		count += lens[i] / 2
	}
	return count
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
