// 剑指 Offer II 097. 子序列的数目
//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
//
//字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
//
//题目数据保证答案符合 32 位带符号整数范围。
//
//
//
//示例 1：
//
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit
//示例 2：
//
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag
//
//
//提示：
//
//0 <= s.length, t.length <= 1000
//s 和 t 由英文字母组成
//
//
//注意：本题与主站 115 题相同： https://leetcode-cn.com/problems/distinct-subsequences/
//
//通过次数1,888提交次数3,494
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/19/21 2:17 PM
package _1dk04

// 记忆化回溯
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	cache := map[int]int{}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if m-i < n-j {
			return 0
		}
		key := i*n + j
		if count, ok := cache[key]; ok {
			return count
		}
		count := dfs(i+1, j)
		if s[i] == t[j] {
			count += dfs(i+1, j+1)
		}
		cache[key] = count
		return count
	}

	return dfs(0, 0)
}

// 动态规划
func numDistinct2(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([]int, n+1)
	dp[n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := max(0, n-m+i); j < n; j++ {
			if s[i] == t[j] {
				dp[j] += dp[j+1]
			}
		}
	}
	return dp[0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
