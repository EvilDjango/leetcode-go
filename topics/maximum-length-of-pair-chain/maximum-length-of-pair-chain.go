// 646. 最长数对链
//给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
//
//现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
//
//给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
//
//
//
//示例：
//
//输入：[[1,2], [2,3], [3,4]]
//输出：2
//解释：最长的数对链是 [1,2] -> [3,4]
//
//
//提示：
//
//给出数对的个数在 [1, 1000] 范围内。
//通过次数23,228提交次数40,132
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/16/21 7:15 PM
package maximum_length_of_pair_chain

import (
	"math"
	"sort"
)

// 动态规划
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	n := len(pairs)
	// dp[i]表示以i位置元素结尾的数对链的最大长度
	dp := make([]int, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		maxPrecursors := 0
		for j := 0; j < i && pairs[j][0] < pairs[i][0]; j++ {
			if pairs[j][1] < pairs[i][0] {
				maxPrecursors = max(maxPrecursors, dp[j])
			}
		}
		dp[i] = maxPrecursors + 1
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 贪心，按照右端点排序，那么右端点最小的应该作为第一个元素，依次类推
func findLongestChain2(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	maxLen := 0
	right := math.MinInt32
	for _, pair := range pairs {
		if pair[0] > right {
			right = pair[1]
			maxLen++
		}
	}
	return maxLen
}
