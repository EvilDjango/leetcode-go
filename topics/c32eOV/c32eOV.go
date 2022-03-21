// 剑指 Offer II 022. 链表中环的入口节点
//给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
//
//说明：不允许修改给定的链表。
//
//
//
//示例 1：
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//示例 2：
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//示例 3：
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//
//
//提示：
//
//链表中节点的数目范围在范围 [0, 104] 内
//-105 <= Node.val <= 105
//pos 的值为 -1 或者链表中的一个有效索引
//
//
//进阶：是否可以使用 O(1) 空间解决此题？
//
//
//
//注意：本题与主站 142 题相同： https://leetcode-cn.com/problems/linked-list-cycle-ii/
//
//通过次数20,507提交次数37,259
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/21 上午11:18
package c32eOV

import . "leetcode-go/container/linklist"

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
