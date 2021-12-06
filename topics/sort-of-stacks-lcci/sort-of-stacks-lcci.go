// 面试题 03.05. 栈排序
//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
//
//示例1:
//
// 输入：
//["SortedStack", "push", "push", "peek", "pop", "peek"]
//[[], [1], [2], [], [], []]
// 输出：
//[null,null,null,1,null,2]
//示例2:
//
// 输入：
//["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
//[[], [], [], [1], [], []]
// 输出：
//[null,null,null,null,null,true]
//说明:
//
//栈中的元素数目在[0, 5000]范围内。
//通过次数15,660提交次数28,980
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/6/21 1:38 PM
package sort_of_stacks_lcci

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack) Peek() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type SortedStack struct {
	s1, s2 *Stack
}

func Constructor() SortedStack {
	return SortedStack{&Stack{}, &Stack{}}
}

func (this *SortedStack) Push(val int) {
	s1, s2 := this.s1, this.s2
	for s2.Len() > 0 && s2.Peek() > val {
		s1.Push(s2.Pop())
	}
	for s1.Len() > 0 && s1.Peek() < val {
		s2.Push(s1.Pop())
	}
	(*s1).Push(val)
}

func (this *SortedStack) Pop() {
	this.drainToS1()
	this.s1.Pop()
}

func (this *SortedStack) Peek() int {
	this.drainToS1()
	return this.s1.Peek()
}

func (this *SortedStack) IsEmpty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

func (this *SortedStack) drainToS1() {
	for this.s2.Len() > 0 {
		this.s1.Push(this.s2.Pop())
	}
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
