// 剑指 Offer II 074. 合并区间
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2：
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//提示：
//
//1 <= intervals.length <= 104
//intervals[i].length == 2
//0 <= starti <= endi <= 104
//
//
//注意：本题与主站 56 题相同： https://leetcode-cn.com/problems/merge-intervals/
//
//通过次数4,616提交次数8,078
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/27/21 11:27 PM
package SsGoHC

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		l, r := intervals[i][0], intervals[i][1]
		if l > right {
			ans = append(ans, []int{left, right})
			left, right = l, r
		} else if r > right {
			right = r
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
