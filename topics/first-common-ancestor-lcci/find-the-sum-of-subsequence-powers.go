// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/23 下午10:30
package first_common_ancestor_lcci

import (
	"math"
	"sort"
)

func sumOfPowers(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([][]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]map[int]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make(map[int]int)
		}
	}
	ans := 0
	mod := 1000_000_007
	for i := 0; i < n; i++ {
		dp[i][1][math.MaxInt32] = 1
		for length := 2; length <= k; length++ {
			for j := 0; j < i; j++ {
				diff := nums[i] - nums[j]
				for v, cnt := range dp[j][length-1] {
					minDiff := min(diff, v)
					dp[i][length][minDiff] = (dp[i][length][minDiff] + cnt) % mod
				}
			}
		}
		for v, cnt := range dp[i][k] {
			ans += v * cnt % mod
			ans %= mod
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
