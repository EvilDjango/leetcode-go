package leetcode_go

import (
	"container/heap"
	"leetcode-go/common"
)

/*
264. 丑数 II
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。


提示：

1 <= n <= 1690
通过次数89,069提交次数155,214

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 6:46 PM
*/

// 优先队列（堆）
func nthUglyNumber(n int) int {
	ints := common.Ints(make([]int, 0, 2*n))
	h := &ints
	heap.Init(h)
	heap.Push(h, 1)
	factors := []int{2, 3, 5}
	seen := make(map[int]bool)
	for i := 1; i < n; i++ {
		cur := heap.Pop(h).(int)
		for _, factor := range factors {
			x := cur * factor
			if !seen[x] {
				heap.Push(h, x)
				seen[x] = true
			}
		}
	}
	return heap.Pop(h).(int)
}

// 动态规划
func nthUglyNumber2(n int) int {
	nums := make([]int, 0, n)
	nums[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		x2, x3, x5 := nums[p2]*2, nums[p3]*3, nums[p5]*5
		min := common.Min(x2, common.Min(x3, x5))
		nums[i] = min
		if x2 == min {
			p2++
		}
		if x3 == min {
			p3++
		}
		if x5 == min {
			p5++
		}
	}
	return nums[n-1]
}
