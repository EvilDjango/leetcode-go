// 面试题 17.20. 连续中值
//随机产生数字并传递给一个方法。你能否完成这个方法，在每次产生新值时，寻找当前所有值的中间值（中位数）并保存。
//
//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4] 的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
//设计一个支持以下两种操作的数据结构：
//
//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//double findMedian() - 返回目前所有元素的中位数。
//示例：
//
//addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2
//通过次数3,813提交次数6,681
// 同295题
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/14/21 11:04 AM
package topic2336

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	small, large hp
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	s, l := &this.small, &this.large
	if s.Len() == 0 || num <= -s.IntSlice[0] {
		heap.Push(s, -num)
		if s.Len() > l.Len()+1 {
			heap.Push(l, -heap.Pop(s).(int))
		}
	} else {
		heap.Push(l, num)
		if l.Len() > s.Len() {
			heap.Push(s, -heap.Pop(l).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	s, l := &this.small, &this.large
	if s.Len() == l.Len() {
		return float64(l.IntSlice[0]-s.IntSlice[0]) / 2
	}
	return float64(-s.IntSlice[0])
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
