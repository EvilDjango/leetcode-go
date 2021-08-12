package topic284

/*
284. 顶端迭代器
给定一个迭代器类的接口，接口包含两个方法： next() 和 hasNext()。设计并实现一个支持 peek() 操作的顶端迭代器 -- 其本质就是把原本应由 next() 方法返回的元素 peek() 出来。

示例:

假设迭代器被初始化为列表 [1,2,3]。

调用 next() 返回 1，得到列表中的第一个元素。
现在调用 peek() 返回 2，下一个元素。在此之后调用 next() 仍然返回 2。
最后一次调用 next() 返回 3，末尾元素。在此之后调用 hasNext() 应该返回 false。
进阶：你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？

通过次数9,683提交次数13,226

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/12/21 10:34 PM
*/

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	it *Iterator
	n  []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, nil}
}

func (this *PeekingIterator) hasNext() bool {
	return this.n != nil || this.it.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.n != nil {
		val := this.n[0]
		this.n = nil
		return val
	}
	return this.it.next()
}

func (this *PeekingIterator) peek() int {
	if this.n == nil {
		this.n = []int{this.it.next()}
	}
	return this.n[0]
}
