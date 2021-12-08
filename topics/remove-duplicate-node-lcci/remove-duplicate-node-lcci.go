// 面试题 02.01. 移除重复节点
//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
//示例1:
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//示例2:
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//提示：
//
//链表长度在[0, 20000]范围内。
//链表元素在[0, 20000]范围内。
//进阶：
//
//如果不得使用临时缓冲区，该怎么解决？
//
//通过次数66,577提交次数98,005
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/8/21 10:00 PM
package remove_duplicate_node_lcci

import . "leetcode-go/container/linklist"

func removeDuplicateNodes(head *ListNode) *ListNode {
	seen := map[int]bool{}
	curr := head
	var prev *ListNode
	for curr != nil {
		if seen[curr.Val] {
			prev.Next = curr.Next
		} else {
			seen[curr.Val] = true
			prev = curr
		}
		curr = curr.Next
	}
	return head
}
