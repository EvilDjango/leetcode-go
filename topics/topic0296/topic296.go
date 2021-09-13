package topic0296

import (
	"leetcode-go/common"
	"math"
)

/*
296. 最佳的碰头地点
有一队人（两人或以上）想要在一个地方碰面，他们希望能够最小化他们的总行走距离。

给你一个 2D 网格，其中各个格子内的值要么是 0，要么是 1。

1 表示某个人的家所处的位置。这里，我们将使用 曼哈顿距离 来计算，其中 distance(p1, p2) = |p2.x - p1.x| + |p2.y - p1.y|。

示例：

输入:

1 - 0 - 0 - 0 - 1
|   |   |   |   |
0 - 0 - 0 - 0 - 0
|   |   |   |   |
0 - 0 - 1 - 0 - 0

输出: 6

解析: 给定的三个人分别住在(0,0)，(0,4) 和 (2,2):
     (0,2) 是一个最佳的碰面点，其总行走距离为 2 + 2 + 2 = 6，最小，因此返回 6。
通过次数2,265提交次数3,670

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/20/21 11:32 AM
*/

func minTotalDistance(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	var people [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				people = append(people, []int{i, j})
			}
		}
	}
	min := math.MaxInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dis := 0
			for _, p := range people {
				dis += common.Abs(i-p[0]) + common.Abs(j-p[1])
			}
			min = common.Min(min, dis)
		}
	}
	return min
}

// 分别计算横纵坐标差值的最小值
func minTotalDistance2(grid [][]int) int {
	var rows, cols []int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rows = append(rows, i)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[j][i] == 1 {
				cols = append(cols, i)
			}
		}
	}
	return minDis(rows) + minDis(cols)
}

func minDis(arr []int) int {
	dis := 0
	l, r := 0, len(arr)-1
	for l < r {
		dis += arr[r] - arr[l]
		r--
		l++
	}
	return dis
}
