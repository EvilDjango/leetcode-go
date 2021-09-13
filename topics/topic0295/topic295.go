package topic0295

import (
	"container/heap"
	"leetcode-go/common"
)

/*
295. 数据流的中位数
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例：

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:

如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
通过次数43,734提交次数83,872

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/19/21 1:42 PM
*/
// 使用两个堆
type MedianFinder struct {
	min common.MinHeap
	max common.MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.min, num)
	heap.Push(&this.max, heap.Pop(&this.min))
	for len(this.max) > len(this.min) {
		heap.Push(&this.min, heap.Pop(&this.max))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.min) > len(this.max) {
		return float64(this.min[0])
	} else {
		return float64(this.min[0]+this.max[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
