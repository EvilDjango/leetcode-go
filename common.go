package leetcode_go

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/27/21 6:16 PM
*/

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Ints []int

func (h *Ints) Len() int {
	return len(*h)
}

func (h *Ints) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Ints) Swap(i, j int) {
	h1 := *h
	h1[i], h1[j] = h1[j], h1[i]
}

func (h *Ints) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Ints) Pop() interface{} {
	h1 := *h
	x := h1[len(h1)-1]
	*h = h1[:len(h1)-1]
	return x
}

type Intervals [][]int

func (in Intervals) Len() int {
	return len(in)
}

func (in Intervals) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}

func (in Intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
