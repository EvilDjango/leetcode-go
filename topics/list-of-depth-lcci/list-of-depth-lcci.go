// 面试题 04.03. 特定深度节点链表
//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
//
//
//示例：
//
//输入：[1,2,3,4,5,null,7,8]
//
//        1
//       /  \
//      2    3
//     / \    \
//    4   5    7
//   /
//  8
//
//输出：[[1],[2,3],[4,5,7],[8]]
//通过次数25,098提交次数31,156
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/3/21 1:59 PM
package list_of_depth_lcci

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	var heads, tails []*ListNode
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		node := &ListNode{Val: root.Val}
		if len(heads) == depth {
			heads = append(heads, node)
			tails = append(tails, node)
		} else {
			tails[depth].Next = node
			tails[depth] = node
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(tree, 0)
	return heads
}
