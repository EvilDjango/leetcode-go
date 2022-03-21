// 剑指 Offer II 026. 重排链表
//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln-1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
//
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
//示例 1:
//
//
//
//输入: head = [1,2,3,4]
//输出: [1,4,2,3]
//示例 2:
//
//
//
//输入: head = [1,2,3,4,5]
//输出: [1,5,2,4,3]
//
//
//提示：
//
//链表的长度范围为 [1, 5 * 104]
//1 <= node.val <= 1000
//
//
//注意：本题与主站 143 题相同：https://leetcode-cn.com/problems/reorder-list/
//
//通过次数14,147提交次数21,362
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/20 下午2:21
package LGjMqU

import . "leetcode-go/container/linklist"

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := reverse(slow.Next)
	slow.Next = nil
	for right != nil {
		next := head.Next
		head.Next = right
		right = right.Next
		head.Next.Next = next
		head = next
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
