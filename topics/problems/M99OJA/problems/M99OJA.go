// 剑指 Offer II 086. 分割回文子字符串
//给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。
//
//回文串 是正着读和反着读都一样的字符串。
//
//
//
//示例 1：
//
//输入：s = "google"
//输出：[["g","o","o","g","l","e"],["g","oo","g","l","e"],["goog","l","e"]]
//示例 2：
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//示例 3：
//
//输入：s = "a"
//输出：[["a"]
//
//
//提示：
//
//1 <= s.length <= 16
//s 仅由小写英文字母组成
//
//
//注意：本题与主站 131 题相同： https://leetcode-cn.com/problems/palindrome-partitioning/
//
//通过次数3,818提交次数5,079
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/23/21 9:53 PM
package problems

// 两个动态规划
func partition(s string) [][]string {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			isPalindrome[i][j] = s[i] == s[j] && (length == 2 || isPalindrome[i+1][j-1])
		}
	}
	dp := make([][][]string, n)
	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			dp[i] = append(dp[i], []string{s[:i+1]})
		}
		for j := 1; j <= i; j++ {
			if isPalindrome[j][i] {
				strs := make([][]string, len(dp[j-1]))
				copy(strs, dp[j-1])
				w := s[j : i+1]
				for _, str := range strs {
					dp[i] = append(dp[i], append(str, w))
				}
			}
		}
	}
	return dp[n-1]
}

// 动态规划+回溯
func partition2(s string) [][]string {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			isPalindrome[i][j] = s[i] == s[j] && (length == 2 || isPalindrome[i+1][j-1])
		}
	}
	var ans [][]string
	var dfs func(int, []string)
	dfs = func(i int, strs []string) {
		if i == n {
			ans = append(ans, append([]string{}, strs...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome[i][j] {
				dfs(j+1, append(strs, s[i:j+1]))
			}
		}
	}
	dfs(0, nil)
	return ans
}
