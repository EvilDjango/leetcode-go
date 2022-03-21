// 剑指 Offer II 024. 反转链表
//给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000
//
//
//进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
//
//注意：本题与主站 206 题相同： https://leetcode-cn.com/problems/reverse-linked-list/
//
//通过次数36,088提交次数47,906
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/20 下午8:15
package UHnkqh

import . "leetcode-go/container/linklist"

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// 递归解法
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
