// 309. 最佳买卖股票时机含冷冻期
//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//示例:
//
//输入: [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//通过次数112,561提交次数185,814
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/8/21 10:38 AM
package topic309

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([]int, 3)
	dp[2] = -prices[0]
	for i := 1; i < n; i++ {
		temp := dp[1]
		dp[1] = max(dp[0], dp[1])
		dp[0] = dp[2] + prices[i]
		dp[2] = max(dp[2], temp-prices[i])
	}
	return max(dp[0], max(dp[1], dp[2]))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
