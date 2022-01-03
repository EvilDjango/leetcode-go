// 剑指 Offer II 053. 二叉搜索树中的中序后继
//给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
//
//节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
//
//
//
//示例 1：
//
//
//
//输入：root = [2,1,3], p = 1
//输出：2
//解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
//示例 2：
//
//
//
//输入：root = [5,3,6,2,4,null,null,1], p = 6
//输出：null
//解释：因为给出的节点没有中序后继，所以答案就返回 null 了。
//
//
//提示：
//
//树中节点的数目在范围 [1, 104] 内。
//-105 <= Node.val <= 105
//树中各节点的值均保证唯一。
//
//
//注意：本题与主站 285 题相同： https://leetcode-cn.com/problems/inorder-successor-in-bst/
//
//通过次数5,996提交次数9,320
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/2/22 4:27 PM
package P5rCT8

import . "leetcode-go/container/tree"

// 直接中序遍历
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var prev *TreeNode
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			if prev == p {
				return root
			}
			stack = stack[:len(stack)-1]
			prev = root
			root = root.Right
		}
	}
	return nil
}

// 利用二叉搜索树的性质来搜索
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var result *TreeNode
	for root != nil {
		if root.Val > p.Val {
			result = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return result
}
