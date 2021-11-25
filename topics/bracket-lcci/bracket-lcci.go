// 面试题 08.09. 括号
//括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
//
//说明：解集不能包含重复的子集。
//
//例如，给出 n = 3，生成结果为：
//
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
//通过次数22,690提交次数27,841
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/24/21 8:03 PM
package bracket_lcci

// 回溯
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var res []string
	backtrack(&res, "", n, 0, 0)
	return res
}

func backtrack(res *[]string, curr string, total, left, right int) {
	if left > total || right > left {
		return
	}
	if left == total && right == total {
		*res = append(*res, curr)
		return
	}
	backtrack(res, curr+"(", total, left+1, right)
	backtrack(res, curr+")", total, left, right+1)
}

// 动态规划
func generateParenthesis2(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-1-j] {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}
		}
	}
	return dp[n]
}

// 解法二的递归版
func generateParenthesis3(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var res []string
	for i := 0; i < n; i++ {
		for _, left := range generateParenthesis3(i) {
			for _, right := range generateParenthesis3(n - 1 - i) {
				res = append(res, "("+left+")"+right)
			}
		}
	}
	return res
}
