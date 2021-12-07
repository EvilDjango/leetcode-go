// 面试题 02.06. 回文链表
//编写一个函数，检查输入的链表是否是回文的。
//
//
//
//示例 1：
//
//输入： 1->2
//输出： false
//示例 2：
//
//输入： 1->2->2->1
//输出： true
//
//
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
//通过次数41,884提交次数86,443
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/7/21 8:00 PM
package palindrome_linked_list_lcci

import . "leetcode-go/container/linklist"

//  O(n) 空间复杂度解法，利用数组
func isPalindrome(head *ListNode) bool {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i, v := range values[:len(values)/2] {
		if v != values[len(values)-i-1] {
			return false
		}
	}
	return true
}

// 递归
func isPalindrome2(head *ListNode) bool {
	var dfs func(head *ListNode) bool
	dfs = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Next) || head.Val != node.Val {
			return false
		}
		head = head.Next
		return true
	}
	return dfs(head)
}

//  O(1) 空间复杂度解法，快慢指针，反转链表
func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head.Next
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
		fast = fast.Next.Next
	}
	left, right := slow, slow
	// 长度为偶数数
	if fast != nil {
		right = right.Next
	}
	palindrome := true
	for left != nil {
		if left.Val != right.Val {
			palindrome = false
		}
		if prev != nil {
			next := prev.Next
			prev.Next = left
			left = prev
			prev = next
		} else {
			left = prev
		}
		right = right.Next
	}
	return palindrome
}
