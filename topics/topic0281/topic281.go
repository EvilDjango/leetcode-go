package topic0281

/*
281. 锯齿迭代器
给出两个一维的向量，请你实现一个迭代器，交替返回它们中间的元素。

示例:

输入:
v1 = [1,2]
v2 = [3,4,5,6]

输出: [1,3,2,4,5,6]

解析: 通过连续调用 next 函数直到 hasNext 函数返回 false，
     next 函数返回值的次序应依次为: [1,3,2,4,5,6]。
拓展：假如给你 k 个一维向量呢？你的代码在这种情况下的扩展性又会如何呢?

拓展声明：
 “锯齿” 顺序对于 k > 2 的情况定义可能会有些歧义。所以，假如你觉得 “锯齿” 这个表述不妥，也可以认为这是一种 “循环”。例如：

输入:
[1,2,3]
[4,5,6,7]
[8,9]

输出: [1,4,8,2,5,9,3,6,7].
通过次数2,324提交次数3,076

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/11/21 2:13 PM
*/

type ZigzagIterator struct {
	v1     []int
	v2     []int
	p1, p2 int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	if len(v1) == 0 {
		v1, v2 = v2, v1
	}
	return &ZigzagIterator{v1, v2, 0, 0}
}

func (this *ZigzagIterator) next() int {
	if !this.hasNext() {
		panic("no more element")
	}
	val := this.v1[this.p1]
	this.p1++
	if this.p2 < len(this.v2) {
		this.p1, this.p2, this.v1, this.v2 = this.p2, this.p1, this.v2, this.v1
	}
	return val
}

func (this *ZigzagIterator) hasNext() bool {
	return this.p1 < len(this.v1)
}
