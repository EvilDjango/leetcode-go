// 面试题 03.02. 栈的最小值
//请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。
//
//
//示例：
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//通过次数27,250提交次数44,429
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/6/21 4:13 PM
package min_stack_lcci

type MinStack struct {
	nums, mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	min := x
	if length := len(this.mins); length > 0 && this.mins[length-1] < x {
		min = this.mins[length-1]
	}
	this.mins = append(this.mins, min)
}

func (this *MinStack) Pop() {
	index := len(this.nums) - 1
	if index < 0 {
		return
	}
	this.nums = this.nums[:index]
	this.mins = this.mins[:index]
}

func (this *MinStack) Top() int {
	return this.nums[(len(this.nums) - 1)]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
