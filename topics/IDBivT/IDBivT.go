// 剑指 Offer II 085. 生成匹配的括号
//正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//示例 2：
//
//输入：n = 1
//输出：["()"]
//
//
//提示：
//
//1 <= n <= 8
//
//
//注意：本题与主站 22 题相同： https://leetcode-cn.com/problems/generate-parentheses/
//
//通过次数5,573提交次数6,550
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/23/21 11:03 PM
package IDBivT

func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}

		}
	}
	return dp[n]
}
