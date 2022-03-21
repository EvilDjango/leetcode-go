// 剑指 Offer II 021. 删除链表的倒数第 n 个结点
//给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
//示例 1：
//
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz
//
//
//进阶：能尝试使用一趟扫描实现吗？
//
//
//
//注意：本题与主站 19 题相同： https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
//
//通过次数17,807提交次数33,760
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/21 下午12:09
package SLwz0R

import . "leetcode-go/container/linklist"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}
	target := size - n
	if target == 0 {
		return head.Next
	}
	curr = head
	for {
		target--
		if target == 0 {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return head
}

// 递归解法
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var dfs func(node, prev *ListNode)
	dfs = func(node, prev *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next, node)
		if n == 1 {
			if prev == nil {
				head = head.Next
			} else {
				prev.Next = node.Next
			}
		}
		n--
	}
	dfs(head, nil)
	return head
}

// 双指针法
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
