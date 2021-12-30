// 剑指 Offer II 061. 和最小的 k 个数对
//给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
//定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
//请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
//
//
//
//示例 1:
//
//输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//示例 2:
//
//输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//示例 3:
//
//输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//提示:
//
//1 <= nums1.length, nums2.length <= 104
//-109 <= nums1[i], nums2[i] <= 109
//nums1, nums2 均为升序排列
//1 <= k <= 1000
//
//
//注意：本题与主站 373 题相同：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/
//
//通过次数3,479提交次数6,384
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/30/21 12:30 PM
package qn8gGX

import "container/heap"

type hp struct {
	a, b  []int
	pairs [][]int
}

func (h *hp) Len() int {
	return len(h.pairs)
}

func (h *hp) Less(i, j int) bool {
	p1, p2 := h.pairs[i], h.pairs[j]
	return h.a[p1[0]]+h.b[p1[1]] < h.a[p2[0]]+h.b[p2[1]]
}

func (h *hp) Swap(i, j int) {
	h.pairs[i], h.pairs[j] = h.pairs[j], h.pairs[i]
}

func (h *hp) Push(v interface{}) {
	h.pairs = append(h.pairs, v.([]int))
}

func (h *hp) Pop() interface{} {
	size := len(h.pairs)
	v := h.pairs[size-1]
	h.pairs = h.pairs[:size-1]
	return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	h := &hp{nums1, nums2, nil}
	for i := 0; i < k && i < m; i++ {
		heap.Push(h, []int{i, 0})
	}
	ans := make([][]int, 0, k)
	for len(ans) < k && h.Len() > 0 {
		ids := heap.Pop(h).([]int)
		ans = append(ans, []int{nums1[ids[0]], nums2[ids[1]]})
		if ids[1]+1 < n {
			heap.Push(h, []int{ids[0], ids[1] + 1})
		}
	}
	return ans
}
