// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/7/21 3:43 PM
package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var nums []int
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return fmt.Sprint(nums)
}

func New(nums ...int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, num := range nums {
		node := &ListNode{Val: num}
		curr.Next = node
		curr = node
	}
	return dummy.Next
}

func CreateCircle(circleStart int, nums ...int) *ListNode {
	if circleStart >= len(nums) {
		panic(fmt.Sprintf("index(%v) out of bound(%v)", circleStart, len(nums)))
	}
	dummy := &ListNode{}
	curr := dummy
	for _, num := range nums {
		node := &ListNode{Val: num}
		curr.Next = node
		curr = node
	}
	tail := curr
	curr = dummy
	for i := 0; i <= circleStart; i++ {
		curr = curr.Next
	}
	tail.Next = curr
	return dummy.Next
}
