// 剑指 Offer II 078. 合并排序链表
//给定一个链表数组，每个链表都已经按升序排列。
//
//请将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//
//输入：lists = []
//输出：[]
//示例 3：
//
//输入：lists = [[]]
//输出：[]
//
//
//提示：
//
//k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4
//
//
//注意：本题与主站 23 题相同： https://leetcode-cn.com/problems/merge-k-sorted-lists/
//
//通过次数5,259提交次数8,226
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 7:40 PM
package vvXgSW

import (
	"container/heap"
	. "leetcode-go/container/linklist"
)

type hp []*ListNode

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}
func (h *hp) Swap(i, j int) {
	a := *h
	a[i], a[j] = a[j], a[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

// 使用优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	dummy := &ListNode{}
	h := &hp{}
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}
	heap.Init(h)
	curr := dummy
	for h.Len() > 0 {
		l := heap.Pop(h).(*ListNode)
		curr.Next = l
		curr = l
		if l.Next != nil {
			heap.Push(h, l.Next)
		}
	}
	return dummy.Next
}

// 归并排序
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var merge func(int, int) *ListNode
	merge = func(left int, right int) *ListNode {
		if left == right {
			return lists[left]
		}
		mid := (right-left)/2 + left
		return mergeTwoLists(merge(left, mid), merge(mid+1, right))
	}
	return merge(0, len(lists)-1)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}
