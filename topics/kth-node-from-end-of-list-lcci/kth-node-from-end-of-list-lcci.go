// 面试题 02.02. 返回倒数第 k 个节点
//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//
//注意：本题相对原题稍作改动
//
//示例：
//
//输入： 1->2->3->4->5 和 k = 2
//输出： 4
//说明：
//
//给定的 k 保证是有效的。
//
//通过次数68,004提交次数87,174
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/8/21 9:57 PM
package kth_node_from_end_of_list_lcci

import . "leetcode-go/container/linklist"

func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}
