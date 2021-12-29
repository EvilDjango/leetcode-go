// 647. 回文子串
//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
//回文字符串 是正着读和倒过来读一样的字符串。
//
//子字符串 是字符串中的由连续字符组成的一个序列。
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
//通过次数137,982提交次数209,354
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/29/21 9:55 PM
package palindromic_substrings

// 动态规划
func countSubstrings(s string) int {
	n := len(s)
	ans := n
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for length := 2; length <= n; length++ {
		for i, j := 0, length-1; j < n; i, j = i+1, j+1 {
			dp[i][j] = s[i] == s[j] && (length == 2 || dp[i+1][j-1])
			if dp[i][j] {
				ans++
			}
		}
	}
	return ans
}

// 中心扩展
func countSubstrings2(s string) int {
	ans := 0
	for i := 0; i < 2*len(s)-1; i++ {
		left, right := i/2, i/2+(1&i)
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			ans++
		}
	}
	return ans
}

// 马拉车算法
func countSubstrings3(s string) int {
	bytes := make([]byte, 0, 2*len(s)+3)
	bytes = append(bytes, '$')
	for _, b := range s {
		bytes = append(bytes, '#', byte(b))
	}
	bytes = append(bytes, '#', '!')
	pivot, right := -1, -1
	lens := make([]int, len(bytes))
	ans := 0
	for i := 1; i < len(bytes)-1; i++ {
		if i < right {
			lens[i] = min(lens[2*pivot-i], right-i+1)
		} else {
			lens[i] = 1
		}
		for bytes[i-lens[i]] == bytes[i+lens[i]] {
			lens[i]++
		}
		if lens[i]+i-1 > right {
			right = lens[i] + i - 1
			pivot = i
		}
		ans += lens[i] / 2
	}
	return ans
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
