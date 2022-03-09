// 面试题19. 正则表达式匹配
//请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。
//
//示例 1:
//
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//示例 2:
//
//输入:
//s = "aa"
//p = "a*"
//输出: true
//解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//示例 3:
//
//输入:
//s = "ab"
//p = ".*"
//输出: true
//解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//示例 4:
//
//输入:
//s = "aab"
//p = "c*a*b"
//输出: true
//解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
//示例 5:
//
//输入:
//s = "mississippi"
//p = "mis*is*p*."
//输出: false
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。
//注意：本题与主站 10 题相同：https://leetcode-cn.com/problems/regular-expression-matching/
//
//通过次数79,366提交次数206,827
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/9 上午10:52
package zheng_ze_biao_da_shi_pi_pei_lcof

// 自己写的版本，动态规划
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.') {
				dp[i][j] = true
				continue
			}
			if p[j-1] != '*' {
				continue
			}
			c := p[j-2]
			// 匹配0个字符
			if dp[i][j-2] {
				dp[i][j] = true
				continue
			}
			for k := i - 1; k >= 0 && (c == '.' || s[k] == c); k-- {
				if dp[k][j-2] {
					dp[i][j] = true
					break
				}
			}
		}
	}
	return dp[m][n]
}

// 参考官方题解写的优化版动态规划
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	matches := func(i, j int) bool {
		return i >= 0 && (p[j] == '.' || p[j] == s[i])
	}
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if matches(i-1, j-2) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i-1, j-1) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
