// 剑指 Offer II 025. 链表中的两数相加
//给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
//可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
//示例1：
//
//
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
//示例2：
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
//示例3：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
//提示：
//
//链表的长度范围为 [1, 100]
//0 <= node.val <= 9
//输入数据保证链表代表的数字无前导 0
//
//
//进阶：如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。
//
//
//
//注意：本题与主站 445 题相同：https://leetcode-cn.com/problems/add-two-numbers-ii/
//
//通过次数13,469提交次数22,471
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/20 下午2:37
package lMSNwu

import . "leetcode-go/container/linklist"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	add := false
	var tail *ListNode
	for l1 != nil || l2 != nil || add {
		val := 0
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		if add {
			val += 1
		}
		node := &ListNode{val % 10, tail}
		add = val >= 10
		tail = node
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return tail
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
