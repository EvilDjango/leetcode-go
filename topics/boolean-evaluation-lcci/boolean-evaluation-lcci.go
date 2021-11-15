// 面试题 08.14. 布尔运算
//给定一个布尔表达式和一个期望的布尔结果 result，布尔表达式由 0 (false)、1 (true)、& (AND)、 | (OR) 和 ^ (XOR) 符号组成。实现一个函数，算出有几种可使该表达式得出 result 值的括号方法。
//
//示例 1:
//
//输入: s = "1^0|0|1", result = 0
//
//输出: 2
//解释: 两种可能的括号方法是
//1^(0|(0|1))
//1^((0|0)|1)
//示例 2:
//
//输入: s = "0&0&0&1^1|0", result = 1
//
//输出: 10
//提示：
//
//运算符的数量不超过 19 个
//通过次数4,356提交次数8,134
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/15/21 1:02 PM
package boolean_evaluation_lcci

// 动态规划
// 使用一个三维数组dp，dp[i][j][k]表示子串s[i:j+1]通过添加括号最终得到结果为k的方案数。
// 注意i和j位置都应该是数字，k只能取0或1
// 要计算dp[i][j]我们只需要考虑从子串s[i:j+1]有多少种拆分为两部分的方案。通过遍历(i,j)范围内的操作符即可。
// 记操作符下标为k，对于每一个可能的k，我们把前半段 dp[i][k-1] 和后半段dp[k+1][j]的结果按情况添加到dp[i][j]当中
// 如果我们从前到后依次遍历的话，就会发现计算过程会依赖还没计算出来的值dp[k+1][j]，行不通。
// 那我们换一个思路，按照子串长度从小到大的顺序来遍历。
func countEval(s string, result int) int {
	n := len(s)
	dp := make([][][2]int, n)
	// 初始化数组
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
	}

	// 子串长度为1的情况
	for i := 0; i < n; i += 2 {
		dp[i][i][s[i]-'0']++
	}
	// 按照子串长度从小到大的顺序依次遍历
	for step := 2; step < n; step += 2 {
		for i := 0; i+step < n; i += 2 {
			for j := i + 1; j < i+step; j += 2 {
				dp[i][i+step][cal(0, 0, s[j])] += dp[i][j-1][0] * dp[j+1][i+step][0]
				dp[i][i+step][cal(0, 1, s[j])] += dp[i][j-1][0]*dp[j+1][i+step][1] + dp[i][j-1][1]*dp[j+1][i+step][0]
				dp[i][i+step][cal(1, 1, s[j])] += dp[i][j-1][1] * dp[j+1][i+step][1]
			}
		}
	}
	return dp[0][n-1][result]
}
func cal(x, y int, op byte) int {
	switch op {
	case '&':
		return x & y
	case '|':
		return x | y
	case '^':
		return x ^ y
	}
	return 0
}

// 记忆化回溯
func countEval2(s string, result int) int {
	n := len(s)
	count := make([][][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([][]int, n)
	}
	return backtrack(s, 0, n-1, count)[result]
}

func backtrack(s string, l, r int, count [][][]int) []int {
	if count[l][r] != nil {
		return count[l][r]
	}
	res := make([]int, 2)
	if l == r {
		res[s[l]-'0']++
		return res
	}
	for i := l + 1; i < r; i += 2 {
		lRes := backtrack(s, l, i-1, count)
		rRes := backtrack(s, i+1, r, count)
		res[cal(0, 0, s[i])] += lRes[0] * rRes[0]
		res[cal(0, 1, s[i])] += lRes[0]*rRes[1] + lRes[1]*rRes[0]
		res[cal(1, 1, s[i])] += lRes[1] * rRes[1]
	}
	count[l][r] = res
	return res
}
