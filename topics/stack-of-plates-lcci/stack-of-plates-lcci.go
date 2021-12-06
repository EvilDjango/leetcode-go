// 面试题 03.03. 堆盘子
//堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。
//
//当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.
//
//示例1:
//
// 输入：
//["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
//[[1], [1], [2], [1], [], []]
// 输出：
//[null, null, null, 2, 1, -1]
//示例2:
//
// 输入：
//["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
//[[2], [1], [2], [3], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, 3]
//通过次数9,067提交次数23,409
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/6/21 3:17 PM
package stack_of_plates_lcci

type StackOfPlates struct {
	cap, size int
	stacks    [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{cap: cap}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	if index := len(this.stacks) - 1; index == -1 || len(this.stacks[index]) == this.cap {
		this.stacks = append(this.stacks, []int{val})
	} else {
		this.stacks[index] = append(this.stacks[index], val)
	}
	this.size++
}

func (this *StackOfPlates) Pop() int {
	if this.size == 0 {
		return -1
	}
	index := len(this.stacks) - 1
	length := len(this.stacks[index])
	val := this.stacks[index][length-1]
	if length > 1 {
		this.stacks[index] = this.stacks[index][:length-1]
	} else {
		this.stacks = this.stacks[:index]
	}
	this.size--
	return val
}

func (this *StackOfPlates) PopAt(index int) int {
	if index >= len(this.stacks) {
		return -1
	}
	length := len(this.stacks[index])
	val := this.stacks[index][length-1]
	if length > 1 {
		this.stacks[index] = this.stacks[index][:length-1]
	} else {
		this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
	}
	this.size--
	return val
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
