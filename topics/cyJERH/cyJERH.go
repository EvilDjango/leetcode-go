// 剑指 Offer II 092. 翻转字符
//如果一个由 '0' 和 '1' 组成的字符串，是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有 '1'）的形式组成的，那么该字符串是 单调递增 的。
//
//我们给出一个由字符 '0' 和 '1' 组成的字符串 s，我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。
//
//返回使 s 单调递增 的最小翻转次数。
//
//
//
//示例 1：
//
//输入：s = "00110"
//输出：1
//解释：我们翻转最后一位得到 00111.
//示例 2：
//
//输入：s = "010110"
//输出：2
//解释：我们翻转得到 011111，或者是 000111。
//示例 3：
//
//输入：s = "00011000"
//输出：2
//解释：我们翻转得到 00000000。
//
//
//提示：
//
//1 <= s.length <= 20000
//s 中只包含字符 '0' 和 '1'
//
//
//注意：本题与主站 926 题相同： https://leetcode-cn.com/problems/flip-string-to-monotone-increasing/
//
//通过次数3,253提交次数4,793
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/23/21 2:59 PM
package cyJERH

// 分别计算前缀1和后缀0的数量
func minFlipsMonoIncr(s string) int {
	n := len(s)
	// prefixOne[i]表示s[:i]中1出现的次数
	prefixOne := make([]int, n+1)
	// suffixZero[i]表示s[i:]中0出现的次数
	suffixZero := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixOne[i] = prefixOne[i-1]
		if s[i-1] == '1' {
			prefixOne[i]++
		}
	}
	for i := n - 1; i >= 0; i-- {
		suffixZero[i] = suffixZero[i+1]
		if s[i] == '0' {
			suffixZero[i]++
		}
	}
	ans := n
	// 枚举1的起始位置
	for i := 0; i <= n; i++ {
		steps := prefixOne[i] + suffixZero[i]
		if steps < ans {
			ans = steps
		}
	}
	return ans
}

// 动态规划
func minFlipsMonoIncr2(s string) int {
	n := len(s)
	// dp[i][0]表示s[:i+1]以0结尾并且满足递增需要的最少转换次数
	// dp[i][1]表示s[:i+1]以1结尾并且满足递增需要的最少转换次数
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{}
	}
	if s[0] == '0' {
		dp[0][1] = 1
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

// 动态规划,滚动变量
func minFlipsMonoIncr3(s string) int {
	zero, one := 0, 0
	if s[0] == '0' {
		one++
	} else {
		zero++
	}
	for i := 1; i < len(s); i++ {
		one = min(one, zero)
		if s[i] == '0' {
			one++
		} else {
			zero++
		}
	}
	return min(zero, one)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
