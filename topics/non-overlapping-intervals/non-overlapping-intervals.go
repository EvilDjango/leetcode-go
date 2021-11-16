// 435. 无重叠区间
//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//
//注意:
//
//可以认为区间的终点总是大于它的起点。
//区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
//示例 1:
//
//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
//示例 2:
//
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//示例 3:
//
//输入: [ [1,2], [2,3] ]
//
//输出: 0
//
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//通过次数103,320提交次数203,794
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/16/21 11:56 AM
package non_overlapping_intervals

import "sort"

// 深度搜索+记忆化 会超时
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	// maxCounts[i]表示从intervals[i]开始的最长无重叠区间序列长度
	maxCounts := make([]int, n)
	maxCount := 0
	for i := 0; i < n; i++ {
		maxCount = max(maxCount, dfs(intervals, maxCounts, i))
	}
	return n - maxCount
}

func dfs(intervals [][]int, counts []int, i int) int {
	if counts[i] > 0 {
		return counts[i]
	}
	successors := 0
	for j := len(intervals) - 1; j > i && intervals[j][0] >= intervals[i][1]; j-- {
		successors = max(successors, dfs(intervals, counts, j))
	}
	counts[i] = successors + 1
	return successors + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 动态规划
func eraseOverlapIntervals2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	// dp[i]表示以i结束的最长区间序列
	dp := make([]int, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		precursors := 0
		for j := 0; j < i && intervals[j][0] < intervals[i][0]; j++ {
			if intervals[j][1] <= intervals[i][0] {
				precursors = max(precursors, dp[j])
			}
		}
		dp[i] = precursors + 1
		maxLen = max(maxLen, dp[i])
	}
	return n - maxLen
}

// 贪心+二分搜索
func eraseOverlapIntervals3(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	n := len(intervals)
	toRemove := 0
	right := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < right {
			toRemove++
		} else {
			right = intervals[i][1]
		}
	}
	return toRemove
}
