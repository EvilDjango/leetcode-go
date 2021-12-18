// 剑指 Offer II 103. 最少的硬币数目
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//你可以认为每种硬币的数量是无限的。
//
//
//
//示例 1：
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//示例 2：
//
//输入：coins = [2], amount = 3
//输出：-1
//示例 3：
//
//输入：coins = [1], amount = 0
//输出：0
//示例 4：
//
//输入：coins = [1], amount = 1
//输出：1
//示例 5：
//
//输入：coins = [1], amount = 2
//输出：2
//
//
//提示：
//
//1 <= coins.length <= 12
//1 <= coins[i] <= 231 - 1
//0 <= amount <= 104
//
//
//注意：本题与主站 322 题相同： https://leetcode-cn.com/problems/coin-change/
//
//通过次数4,695提交次数9,484
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/18/21 7:01 PM
package gaM7Ch

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > amount {
			break
		}
		dp[coin] = 1
	}
	for i := 2; i <= amount; i++ {
		if dp[i] != 0 {
			continue
		}
		min := math.MaxInt64
		for _, coin := range coins {
			if coin >= i {
				break
			}
			if dp[i-coin] != 0 && dp[i-coin]+1 < min {
				min = dp[i-coin] + 1
			}
		}
		if min != math.MaxInt64 {
			dp[i] = min
		}
	}
	return dp[amount]
}
