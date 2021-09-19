// 面试题 17.09. 第 k 个数
//有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
//
//示例 1:
//
//输入: k = 5
//
//输出: 9
//通过次数13,773提交次数25,051
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/19/21 2:44 PM
package get_kth_magic_number_lcci

import (
	"container/heap"
	"sort"
)

// 使用最小堆
func getKthMagicNumber(k int) int {
	h := &hp{[]int{1}}
	prev := -1
	for i := 1; i < k; i++ {
		num := heap.Pop(h).(int)
		if prev == num {
			i--
			continue
		} else {
			prev = num
		}
		heap.Push(h, num*3)
		heap.Push(h, num*5)
		heap.Push(h, num*7)
	}
	for h.IntSlice[0] == prev {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	val := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return val
}

// 三指针法
func getKthMagicNumber2(k int) int {
	i3, i5, i7 := 0, 0, 0
	nums := make([]int, k)
	nums[0] = 1
	for i := 1; i < k; i++ {
		nums[i] = min(min(nums[i3]*3, nums[i5]*5), nums[i7]*7)
		if nums[i] == nums[i3]*3 {
			i3++
		}
		if nums[i] == nums[i5]*5 {
			i5++
		}
		if nums[i] == nums[i7]*7 {
			i7++
		}
	}
	return nums[k-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
