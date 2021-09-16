// 面试题 17.14. 最小K个数
//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
//示例：
//
//输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//提示：
//
//0 <= len(arr) <= 100000
//0 <= k <= min(100000, len(arr))
//通过次数75,332提交次数125,988
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/16/21 6:36 PM
package smallest_k_lcci

import (
	"container/heap"
	"sort"
)

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, i := range arr[k:] {
		if i < h.IntSlice[0] {
			h.IntSlice[0] = i
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(x interface{}) {}
func (h *hp) Pop() interface{}   { return nil }
