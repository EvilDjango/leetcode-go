// 剑指 Offer II 054. 所有大于等于节点的值之和
//给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
//
//
//
//提醒一下，二叉搜索树满足下列约束条件：
//
//节点的左子树仅包含键 小于 节点键的节点。
//节点的右子树仅包含键 大于 节点键的节点。
//左右子树也必须是二叉搜索树。
//
//
//示例 1：
//
//
//
//输入：root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
//示例 2：
//
//输入：root = [0,null,1]
//输出：[1,null,1]
//示例 3：
//
//输入：root = [1,0,2]
//输出：[3,3,2]
//示例 4：
//
//输入：root = [3,2,4,1]
//输出：[7,9,4,10]
//
//
//提示：
//
//树中的节点数介于 0 和 104 之间。
//每个节点的值介于 -104 和 104 之间。
//树中的所有值 互不相同 。
//给定的树为二叉搜索树。
//
//
//注意：
//
//本题与主站 538 题相同： https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
//本题与主站 1038 题相同：https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/
//通过次数4,760提交次数5,557
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/2/22 4:20 PM
package w6cpku

import . "leetcode-go/container/tree"

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Right
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum += curr.Val
			curr.Val = sum
			curr = curr.Left
		}
	}
	return root
}
