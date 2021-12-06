// 面试题 03.01. 三合一
//三合一。描述如何只用一个数组来实现三个栈。
//
//你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。
//
//构造函数会传入一个stackSize参数，代表每个栈的大小。
//
//示例1:
//
// 输入：
//["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
//[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
// 输出：
//[null, null, null, 1, -1, -1, true]
//说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
//示例2:
//
// 输入：
//["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
//[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, -1, -1]
//
//
//提示：
//
//0 <= stackNum <= 2
//通过次数11,959提交次数22,331
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/6/21 4:25 PM
package three_in_one_lcci

type TripleInOne struct {
	size   int
	lens   [3]int
	stacks []int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{size: stackSize, stacks: make([]int, stackSize*3)}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.lens[stackNum] == this.size {
		return
	}
	this.stacks[this.size*stackNum+this.lens[stackNum]] = value
	this.lens[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.lens[stackNum] == 0 {
		return -1
	}
	val := this.stacks[this.size*stackNum+this.lens[stackNum]-1]
	this.lens[stackNum]--
	return val
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.lens[stackNum] == 0 {
		return -1
	}
	val := this.stacks[this.size*stackNum+this.lens[stackNum]-1]
	return val
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.lens[stackNum] == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
