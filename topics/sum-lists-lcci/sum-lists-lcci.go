// 面试题 02.05. 链表求和
//给定两个用链表表示的整数，每个节点包含一个数位。
//
//这些数位是反向存放的，也就是个位排在链表首部。
//
//编写函数对这两个整数求和，并用链表形式返回结果。
//
//
//
//示例：
//
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
//示例：
//
//输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//通过次数39,367提交次数84,315
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/8/21 8:00 PM
package sum_lists_lcci

import . "leetcode-go/container/linklist"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := false
	for l1 != nil || l2 != nil || carry {
		val := 0
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if carry {
			val++
		}
		carry = val > 9
		node := &ListNode{Val: val % 10}
		curr.Next = node
		curr = node
	}
	return dummy.Next
}
