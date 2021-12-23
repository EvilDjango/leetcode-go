// 剑指 Offer II 094. 最少回文分割
//给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。
//
//返回符合要求的 最少分割次数 。
//
//
//
//示例 1：
//
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//示例 2：
//
//输入：s = "a"
//输出：0
//示例 3：
//
//输入：s = "ab"
//输出：1
//
//
//提示：
//
//1 <= s.length <= 2000
//s 仅由小写英文字母组成
//
//
//注意：本题与主站 132 题相同： https://leetcode-cn.com/problems/palindrome-partitioning-ii/
//
//通过次数2,213提交次数3,668
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/21/21 1:40 PM
package omKAoA

// 使用中心扩展法来枚举每个位置的全部回文左端点，然后用dp求解答案
func minCut(s string) int {
	n := len(s)
	lefts := make([][]int, n)
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			lefts[r] = append(lefts[r], l)
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			lefts[r] = append(lefts[r], l)
			l--
			r++
		}
	}
	// dp[i]表示s[:i+1]需要切割的次数
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = i
		for _, left := range lefts[i] {
			if left == 0 {
				dp[i] = 0
				break
			}
			dp[i] = min(dp[i], dp[left-1]+1)
		}
	}
	return dp[n-1]
}

// 使用dp求解回文，然后用dp求解答案
func minCut2(s string) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
	}
	for length := 2; length <= n; length++ {
		for left := 0; left+length-1 < n; left++ {
			right := left + length - 1
			isPalindrome[left][right] = s[left] == s[right] && (length == 2 || isPalindrome[left+1][right-1])
		}
	}
	// dp[i]表示s[:i+1]最少切割数
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = i
		if isPalindrome[0][i] {
			dp[i] = 0
			continue
		}
		for j := 1; j <= i; j++ {
			if isPalindrome[j][i] {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	return dp[n-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
