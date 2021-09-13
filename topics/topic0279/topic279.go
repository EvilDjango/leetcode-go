package topic0279

import (
	"leetcode-go/common"
	"math"
)

/*
279. 完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。



示例 1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9

提示：

1 <= n <= 104
通过次数185,684提交次数295,395

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/11/21 12:31 PM
*/

// 暴力算法

var dict = make(map[int]int)

func numSquares(n int) int {
	if v, ok := dict[n]; ok {
		return v
	}
	if n == 0 {
		return 0
	}
	root := int(math.Sqrt(float64(n)))
	min := math.MaxInt64
	for i := root; i > 0; i-- {
		sub := numSquares(n - i*i)
		min = common.Min(min, sub)
		if sub == 0 || sub == 1 {
			break
		}
	}
	dict[n] = min + 1
	return min + 1
}

// 动态规划
func numSquares2(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		for j := 1; j*j <= i; j++ {
			min = common.Min(min, 1+dp[i-j*j])
		}
		dp[i] = min
	}
	return dp[n]
}

// 一个数学定理可以帮助解决本题：「四平方和定理」。
//
//四平方和定理证明了任意一个正整数都可以被表示为至多四个正整数的平方和。这给出了本题的答案的上界。
//
//同时四平方和定理包含了一个更强的结论：当且仅当 n \neq 4^k \times (8m+7)n
//
//​
// =4
//k
// ×(8m+7) 时，nn 可以被表示为至多三个正整数的平方和。因此，当 n = 4^k \times (8m+7)n=4
//k
// ×(8m+7) 时，nn 只能被表示为四个正整数的平方和。此时我们可以直接返回 44。
//
func numSquares3(n int) int {
	if isFour(n) {
		return 4
	}
	if isSquare(n) {
		return 1
	}
	for i := 1; i*i < n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3

}

func isSquare(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func isFour(n int) bool {
	for n%4 == 0 {
		n /= 4
	}
	return n%8 == 7

}
