// 面试题 03.04. 化栈为队
//实现一个MyQueue类，该类用两个栈来实现一个队列。
//
//
//示例：
//
//MyQueue queue = new MyQueue();
//
//queue.push(1);
//queue.push(2);
//queue.peek();  // 返回 1
//queue.pop();   // 返回 1
//queue.empty(); // 返回 false
//
//说明：
//
//你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size 和 is empty 操作是合法的。
//你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
//通过次数24,581提交次数34,435
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/6/21 2:59 PM
package implement_queue_using_stacks_lcci

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

type MyQueue struct {
	s1, s2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{&Stack{}, &Stack{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s2.Len() == 0 {
		this.drainToS2()
	}
	return this.s2.Pop()

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.s2.Len() == 0 {
		this.drainToS2()
	}
	return this.s2.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

func (this *MyQueue) drainToS2() {
	for this.s1.Len() > 0 {
		this.s2.Push(this.s1.Pop())
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
