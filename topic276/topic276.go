package topic276

/*
276. 栅栏涂色
有 k 种颜色的涂料和一个包含 n 个栅栏柱的栅栏，请你按下述规则为栅栏设计涂色方案：

每个栅栏柱可以用其中 一种 颜色进行上色。
相邻的栅栏柱 最多连续两个 颜色相同。
给你两个整数 k 和 n ，返回所有有效的涂色 方案数 。



示例 1：


输入：n = 3, k = 2
输出：6
解释：所有的可能涂色方案如上图所示。注意，全涂红或者全涂绿的方案属于无效方案，因为相邻的栅栏柱 最多连续两个 颜色相同。
示例 2：

输入：n = 1, k = 1
输出：1
示例 3：

输入：n = 7, k = 2
输出：42


提示：

1 <= n <= 50
1 <= k <= 105
题目数据保证：对于输入的 n 和 k ，其答案在范围 [0, 231 - 1] 内
通过次数6,543提交次数13,574

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/6/21 6:00 PM
*/

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
	}
	for i := 0; i < k; i++ {
		dp[0][i] = 1
		dp[1][i] = k
	}

	for i := 2; i < n; i++ {
		for j := 0; j < k; j++ {
			sum := 0
			for l := 0; l < k; l++ {
				if j == l {
					continue
				}
				sum += dp[i-1][l]
				sum += dp[i-2][l]
			}
			dp[i][j] = sum
		}
	}
	ans := 0
	for i := 0; i < k; i++ {
		ans += dp[n-1][i]
	}
	return ans
}

// 和方法一都是动态规划，但是做了优化
func numWays2(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	if k == 1 {
		return 0
	}

	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < k; i++ {
		dp[i][0] = k - 1
		dp[i][1] = k*k - k
	}
	for i := 2; i < n; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			dp[j][1] += dp[j][0]
			dp[j][0] = dp[j][1] - dp[j][0]
			sum += dp[j][1]
		}
		for j := 0; j < k; j++ {
			dp[j][1] = sum - dp[j][1]
		}
	}
	ans := 0
	for i := 0; i < k; i++ {
		ans += dp[i][1]
	}
	return ans / (k - 1)
}

// 进一步优化
func numWays3(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	if k == 1 {
		return 0
	}
	prev, last := 1, k
	for i := 2; i < n; i++ {
		prev, last = last, (k-1)*(prev+last)
	}
	return k * last
}
