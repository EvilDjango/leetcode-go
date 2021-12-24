// 剑指 Offer II 077. 链表排序
//给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//示例 1：
//
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//示例 2：
//
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目在范围 [0, 5 * 104] 内
//-105 <= Node.val <= 105
//
//
//进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//
//注意：本题与主站 148 题相同：https://leetcode-cn.com/problems/sort-list/
//
//通过次数6,883提交次数11,758
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 9:23 PM
package _WHec2

import . "leetcode-go/container/linklist"

// 冒泡排序,会超时
func sortList2(head *ListNode) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	for i := 0; i < length-1; i++ {
		prev := head
		curr = head.Next
		change := false
		for j := 1; j < length-i; j++ {
			if prev.Val > curr.Val {
				curr.Val, prev.Val = prev.Val, curr.Val
				change = true
			}
			prev, curr = curr, curr.Next
		}
		if !change {
			break
		}
	}
	return head
}

// 自顶向下归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow.Next
	slow.Next = nil
	return mergeTwoList(sortList(head), sortList(right))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
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

// 自底向上的归并排序
func sortList3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}
	dummy := &ListNode{}
	dummy.Next = head
	for subLen := 1; subLen < length; subLen *= 2 {
		index := 0
		prev := dummy
		for index+subLen < length {
			slow, fast := prev, prev
			for i := 0; i < subLen; i++ {
				slow = slow.Next
			}
			for i := 0; i < min(2*subLen, length-index); i++ {
				fast = fast.Next
			}
			l2 := slow.Next
			slow.Next = nil
			next := fast.Next
			fast.Next = nil
			prev = merge(prev, slow, l2, fast, next)
			index += 2 * subLen
		}
	}
	return dummy.Next
}

func merge(prev, l1Tail, l2, l2Tail, next *ListNode) *ListNode {
	l1 := prev.Next
	curr := prev
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
		l1Tail.Next = next
		return l1Tail
	}
	if l2 != nil {
		curr.Next = l2
		l2Tail.Next = next
		return l2Tail
	}
	return nil
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
