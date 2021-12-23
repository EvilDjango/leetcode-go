// 剑指 Offer II 095. 最长公共子序列
//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
//一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
//
//
//示例 1：
//
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace" ，它的长度为 3 。
//示例 2：
//
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc" ，它的长度为 3 。
//示例 3：
//
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0 。
//
//
//提示：
//
//1 <= text1.length, text2.length <= 1000
//text1 和 text2 仅由小写英文字符组成。
//
//
//注意：本题与主站 1143 题相同： https://leetcode-cn.com/problems/longest-common-subsequence/
//
//通过次数8,712提交次数13,385
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/21/21 11:08 AM
package qJnOS7

// 记忆化回溯
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}
	// text1[i:]和text2[j:]的公共子序列最大长度
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m || j == n {
			return 0
		}
		if cache[i][j] >= 0 {
			return cache[i][j]
		}
		var count int
		if text1[i] == text2[j] {
			count = dfs(i+1, j+1) + 1
		} else {
			count = max(dfs(i, j+1), dfs(i+1, j))
		}
		cache[i][j] = count
		return count
	}
	return dfs(0, 0)
}

// 动态规划
func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
