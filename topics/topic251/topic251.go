package topic251

/*
251. 展开二维向量
请设计并实现一个能够展开二维向量的迭代器。该迭代器需要支持 next 和 hasNext 两种操作。



示例：

Vector2D iterator = new Vector2D([[1,2],[3],[4]]);

iterator.next(); // 返回 1
iterator.next(); // 返回 2
iterator.next(); // 返回 3
iterator.hasNext(); // 返回 true
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 4
iterator.hasNext(); // 返回 false


注意：

请记得 重置 在 Vector2D 中声明的类变量（静态变量），因为类变量会 在多个测试用例中保持不变，影响判题准确。请 查阅 这里。
你可以假定 next() 的调用总是合法的，即当 next() 被调用时，二维向量总是存在至少一个后续元素。


进阶：尝试在代码中仅使用 C++ 提供的迭代器 或 Java 提供的迭代器。

通过次数3,121提交次数5,808

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/23/21 2:03 PM
*/

type Vector2D struct {
	vec      [][]int
	row, col int
}

func Constructor(vec [][]int) Vector2D {
	v := Vector2D{vec, 0, -1}
	v.searchForNext()
	return v
}

func (this *Vector2D) Next() int {
	value := this.vec[this.row][this.col]
	this.searchForNext()
	return value
}

func (this *Vector2D) HasNext() bool {
	return this.row < len(this.vec) && this.col < len(this.vec[this.row])
}

func (this *Vector2D) searchForNext() bool {
	for this.row < len(this.vec) {
		if this.col+1 < len(this.vec[this.row]) {
			this.col++
			return true
		} else {
			this.col = -1
			this.row++
		}
	}
	return false
}
