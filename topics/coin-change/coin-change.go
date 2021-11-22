// 322. 零钱兑换
//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
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
//通过次数326,860提交次数731,525
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/22/21 8:18 PM
package coin_change

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = -1
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j+coins[i] <= amount; j++ {
			if dp[j] == -1 {
				continue
			}
			if dp[j+coins[i]] == -1 {
				dp[j+coins[i]] = dp[j] + 1
			} else {
				dp[j+coins[i]] = min(dp[j+coins[i]], dp[j]+1)
			}
		}
		fmt.Println()
	}
	return dp[amount]

}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
