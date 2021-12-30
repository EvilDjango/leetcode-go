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

// 贪心+二分，按左端点排序，然后维护一个数组来表示各个长度的序列右端点的最小值
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	// end[i]表示长度为i+1的序列最右端的最小值
	var ends []int
	for _, pair := range pairs {
		l, r := 0, len(ends)
		for l < r {
			mid := (r-l)>>1 + l
			if ends[mid] >= pair[0] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l == len(ends) {
			ends = append(ends, pair[1])
		} else if pair[1] < ends[l] {
			ends[l] = pair[1]
		}
	}
	return len(ends)
}

// 贪心解法2，按右端点排序，那么最长序列的第一个数对一定是第一个数对，依次类推
func findLongestChain2(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	maxLen := 0
	right := math.MinInt64
	for _, pair := range pairs {
		if pair[0] > right {
			maxLen++
			right = pair[1]
		}
	}
	return maxLen
}

// 动态规划
func findLongestChain3(pairs [][]int) int {
	n := len(pairs)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	// dp[i]表示以pairs[i]结尾的序列的最大长度
	dp := make([]int, n)
	maxLen := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			// 往左遇到的第一个可以拼接的数对一定可以组成最长数对，因为pairs[j]的左端点一定是可拼接的数对中最大的，留给左侧的空间最多
			if pairs[j][1] < pairs[i][0] {
				dp[i] = dp[j] + 1
				maxLen = dp[i]
				break
			}
		}
	}
	return maxLen
}
