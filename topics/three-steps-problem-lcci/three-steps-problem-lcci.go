// 面试题 08.01. 三步问题
//三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
//
//示例1:
//
// 输入：n = 3
// 输出：4
// 说明: 有四种走法
//示例2:
//
// 输入：n = 5
// 输出：13
//提示:
//
//n范围在[1, 1000000]之间
//通过次数44,374提交次数122,124
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/28/21 5:04 PM
package three_steps_problem_lcci

// 记忆化回溯
func waysToStep(n int) int {
	cache := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		cache[i] = -1
	}
	return backtrack(cache, n)
}

func backtrack(cache []int, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if cache[n] != -1 {
		return cache[n]
	}
	res := 0
	res = (res + backtrack(cache, n-3)) % 1000000007
	res = (res + backtrack(cache, n-2)) % 1000000007
	res = (res + backtrack(cache, n-1)) % 1000000007
	cache[n] = res
	return res
}

// 动态规划
func waysToStep2(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}
	return dp[n]
}

// 动态规划，由于dp[i]只跟dp[i-1],dp[i-2],dp[i-3]有关，所以可以只用三个变量来记录前面的结果
func waysToStep3(n int) int {
	if n < 3 {
		return n
	}
	a0, a1, a2 := 1, 1, 2
	for i := 3; i < n+1; i++ {
		a0, a1, a2 = a1, a2, (a0+a1+a2)%1000000007
	}
	return a2
}
