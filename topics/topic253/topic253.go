package topic253

import (
	"container/heap"
	"leetcode-go/common"
	"sort"
)

/*
253. 会议室 II
给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。



示例 1：

输入：intervals = [[0,30],[5,10],[15,20]]
输出：2
示例 2：

输入：intervals = [[7,10],[2,4]]
输出：1


提示：

1 <= intervals.length <= 104
0 <= starti < endi <= 106
通过次数27,618提交次数56,022
Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/25/21 9:27 PM
*/

func minMeetingRooms(intervals [][]int) int {
	h := &common.Ints{}
	heap.Init(h)
	sort.Sort(common.Intervals(intervals))
	for i := 0; i < len(intervals); i++ {
		if h.Len() > 0 {
			x := heap.Pop(h).(int)
			if x > intervals[i][0] {
				heap.Push(h, x)
			}
		}
		heap.Push(h, intervals[i][1])
	}
	return h.Len()
}

func minMeetingRooms2(intervals [][]int) int {
	n := len(intervals)
	starts := common.Ints(make([]int, n))
	ends := common.Ints(make([]int, n))
	for i := 0; i < n; i++ {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}
	sort.Sort(&starts)
	sort.Sort(&ends)
	ans := 0
	for i, j := 0, 0; i < n; i++ {
		if starts[i] < ends[j] {
			ans++
		} else {
			j++
		}
	}
	return ans
}
