package leetcode_go

import "math"

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

func (h Ints) Len() int {
	return len(h)
}

func (h Ints) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Ints) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
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

func UpperBound(nums []int, target int) int {
	return UpperBoundPart(nums, target, 0, len(nums))
}

func UpperBoundPart(nums []int, target, l, r int) int {
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func LowerBound(nums []int, target int) int {
	return LowerBoundPart(nums, target, 0, len(nums))
}

func LowerBoundPart(nums []int, target, l, r int) int {
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func MinWithoutOne(nums []int) []int {
	return extremeWithoutOne(nums, true)
}

func MaxWithoutOne(nums []int) []int {
	return extremeWithoutOne(nums, false)
}

func MinWithoutOne2(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}
	min, second := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num < min {
			second = min
			min = num
		} else if num < second {
			second = num
		}
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] == min {
			ret[i] = second
		} else {
			ret[i] = min
		}
	}
	return ret
}

func MaxWithoutOne2(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}
	max, second := math.MinInt32, math.MinInt32
	for _, num := range nums {
		if num > max {
			second = max
			max = num
		} else if num > second {
			second = num
		}
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] == max {
			ret[i] = second
		} else {
			ret[i] = max
		}
	}
	return ret
}
func extremeWithoutOne(nums []int, min bool) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}
	pre := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < pre[i-1] == min {
			pre[i] = nums[i]
		} else {
			pre[i] = pre[i-1]
		}
	}
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffix[i+1] == min {
			suffix[i] = nums[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		if i-1 >= 0 && i+1 < n {
			if pre[i-1] < suffix[i+1] == min {
				ret[i] = pre[i-1]
			} else {
				ret[i] = suffix[i+1]
			}
		} else if i-1 >= 0 {
			ret[i] = pre[i-1]
		} else if i+1 < n {
			ret[i] = suffix[i+1]
		}
	}
	return ret
}
