// 剑指 Offer II 027. 回文链表
//给定一个链表的 头节点 head ，请判断其是否为回文链表。
//
//如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
//
//
//
//示例 1：
//
//
//
//输入: head = [1,2,3,3,2,1]
//输出: true
//示例 2：
//
//
//
//输入: head = [1,2]
//输出: false
//
//
//提示：
//
//链表 L 的长度范围为 [1, 105]
//0 <= node.val <= 9
//
//
//进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/18 上午10:46
package aMhZSa

import (
	. "leetcode-go/container/linklist"
)

// 先用快慢指针法找到中点，然后反转链表来对比前后两段值是否对应相等
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	left, right := slow, slow.Next
	reverse(nil, head, slow)
	if fast == nil {
		left = slow.Next
	}
	lCur, rCur := left, right
	for lCur != nil && lCur.Val == rCur.Val {
		lCur = lCur.Next
		rCur = rCur.Next
	}
	reverse(right, left, head)
	return lCur == nil
}

func reverse(prev, head, tail *ListNode) {
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}
