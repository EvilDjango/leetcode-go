package topic0265

import (
	"leetcode-go/common"
	"math"
)

/*
265. 粉刷房子 II
假如有一排房子，共 n 个，每个房子可以被粉刷成 k 种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。

当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x k 的矩阵来表示的。

例如，costs[0][0] 表示第 0 号房子粉刷成 0 号颜色的成本花费；costs[1][2] 表示第 1 号房子粉刷成 2 号颜色的成本花费，以此类推。请你计算出粉刷完所有房子最少的花费成本。

注意：

所有花费均为正整数。

示例：

输入: [[1,5,3],[2,9,4]]
输出: 5
解释: 将 0 号房子粉刷成 0 号颜色，1 号房子粉刷成 2 号颜色。最少花费: 1 + 4 = 5;
     或者将 0 号房子粉刷成 2 号颜色，1 号房子粉刷成 0 号颜色。最少花费: 3 + 2 = 5.
进阶：
您能否在 O(nk) 的时间复杂度下解决此问题？

通过次数5,565提交次数9,711


Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 8:17 PM
*/

// 动态规划
func minCostII(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k := len(costs[0])
	if k == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	for i := 0; i < k; i++ {
		dp[0][i] = costs[0][i]
	}
	for i := 1; i < n; i++ {
		minWithoutOne := common.MinWithoutOne(dp[i-1])
		for j := 0; j < k; j++ {
			dp[i][j] = costs[i][j] + minWithoutOne[j]
		}
	}
	ans := math.MaxInt32
	for _, cost := range dp[n-1] {
		ans = common.Min(ans, cost)
	}
	return ans
}

// 动态规划,参考了优秀题解
func minCostII2(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	k := len(costs[0])
	if k < 2 {
		return costs[0][0]
	}
	min, secondMin, minColor := 0, 0, 0
	for _, cost := range costs {
		tMin, tSecond, tMinColor := math.MaxInt32, math.MaxInt32, -1
		for i := 0; i < k; i++ {
			var minCost int
			if i == minColor {
				minCost = secondMin + cost[i]
			} else {
				minCost = min + cost[i]
			}
			if minCost < tMin {
				tSecond = tMin
				tMin = minCost
				tMinColor = i
			} else if minCost < tSecond {
				tSecond = minCost
			}
		}
		min, secondMin, minColor = tMin, tSecond, tMinColor
	}
	return min
}
