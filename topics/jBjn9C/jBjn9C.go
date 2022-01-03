// 剑指 Offer II 059. 数据流的第 K 大数值
//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
//请实现 KthLargest 类：
//
//KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
//int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
//
//
//示例：
//
//输入：
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//输出：
//[null, 4, 5, 5, 8, 8]
//
//解释：
//KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//kthLargest.add(3);   // return 4
//kthLargest.add(5);   // return 5
//kthLargest.add(10);  // return 5
//kthLargest.add(9);   // return 8
//kthLargest.add(4);   // return 8
//
//
//提示：
//
//1 <= k <= 104
//0 <= nums.length <= 104
//-104 <= nums[i] <= 104
//-104 <= val <= 104
//最多调用 add 方法 104 次
//题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
//
//
//注意：本题与主站 703 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
//
//通过次数4,189提交次数6,651
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/31/21 1:34 PM
package jBjn9C

import "container/heap"

type hp []int

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	a := *h
	return a[i] < a[j]
}
func (h *hp) Swap(i, j int) {
	a := *h
	a[i], a[j] = a[j], a[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type KthLargest struct {
	k int
	h *hp
}

func Constructor(k int, nums []int) KthLargest {
	n := k
	if n > len(nums) {
		n = len(nums)
	}
	h := hp(nums[:n])
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		if h[0] < nums[i] {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}
	return KthLargest{k, &h}
}

func (this *KthLargest) Add(val int) int {
	h := this.h
	if h.Len() < this.k {
		heap.Push(this.h, val)
	} else if (*h)[0] < val {
		heap.Pop(h)
		heap.Push(h, val)
	}
	return (*h)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
