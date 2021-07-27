package leetcode_go

import (
	"sort"
)

/*
252. 会议室
给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，请你判断一个人是否能够参加这里面的全部会议。



示例 1：

输入：intervals = [[0,30],[5,10],[15,20]]
输出：false
示例 2：

输入：intervals = [[7,10],[2,4]]
输出：true


提示：

0 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti < endi <= 106
通过次数10,277提交次数18,456


Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/23/21 3:17 PM
*/

func canAttendMeetings(intervals [][]int) bool {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if conflict(intervals[i], intervals[j]) {
				return false
			}
		}
	}
	return true
}

func conflict(a, b []int) bool {
	if a[0] > b[0] {
		a, b = b, a
	}
	return a[1] >= b[0]
}

// 先排序
func canAttendMeetings2(intervals [][]int) bool {
	sort.Sort(Intervals(intervals))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
