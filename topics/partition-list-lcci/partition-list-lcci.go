// 面试题 02.04. 分割链表
//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
//你不需要 保留 每个分区中各节点的初始相对位置。
//
//
//
//示例 1：
//
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//示例 2：
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//
//
//提示：
//
//链表中节点的数目在范围 [0, 200] 内
//-100 <= Node.val <= 100
//-200 <= x <= 200
//通过次数27,036提交次数41,551
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/8/21 9:18 PM
package partition_list_lcci

import . "leetcode-go/container/linklist"

func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	curr1, curr2 := small, large
	for head != nil {
		if head.Val < x {
			curr1.Next = head
			curr1 = head
		} else {
			curr2.Next = head
			curr2 = head
		}
		head = head.Next
	}
	curr2.Next = nil
	curr1.Next = large.Next
	return small.Next
}
