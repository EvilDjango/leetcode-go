// 354. 俄罗斯套娃信封问题
//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
//当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
//请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
//注意：不允许旋转信封。
//
//
//示例 1：
//
//输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出：3
//解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//示例 2：
//
//输入：envelopes = [[1,1],[1,1],[1,1]]
//输出：1
//
//
//提示：
//
//1 <= envelopes.length <= 5000
//envelopes[i].length == 2
//1 <= wi, hi <= 104
//通过次数72,475提交次数158,963
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/17/21 10:33 AM
package russian_doll_envelopes

import "sort"

// 动态规划
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	n := len(envelopes)
	// dp[i]表示以i位置的信封为最大信封的套娃的最大深度
	dp := make([]int, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		// i节点前面比较小的信封序列的最大长度
		maxPres := 0
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				maxPres = max(maxPres, dp[j])
			}
		}
		dp[i] = maxPres + 1
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

// 贪心+二分查找
func maxEnvelopes2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	n := len(envelopes)
	// minH[i]表示长度为i的最长序列末尾的高度
	minH := make([]int, n+1)
	maxLen := 0
	for _, env := range envelopes {
		// 注意，这里查找的是lower bound，即第一个大于等于目标值的元素下标
		l, r := 1, maxLen+1
		for l < r {
			mid := (r-l)>>1 + l
			if minH[mid] < env[1] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		minH[r] = env[1]
		maxLen = max(maxLen, r)
	}
	return maxLen
}
