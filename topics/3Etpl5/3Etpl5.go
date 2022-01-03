// 剑指 Offer II 049. 从根节点到叶节点的路径数字之和
//给定一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//
//每条从根节点到叶节点的路径都代表一个数字：
//
//例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//计算从根节点到叶节点生成的 所有数字之和 。
//
//叶节点 是指没有子节点的节点。
//
//
//
//示例 1：
//
//
//输入：root = [1,2,3]
//输出：25
//解释：
//从根到叶子节点路径 1->2 代表数字 12
//从根到叶子节点路径 1->3 代表数字 13
//因此，数字总和 = 12 + 13 = 25
//示例 2：
//
//
//输入：root = [4,9,0,5,1]
//输出：1026
//解释：
//从根到叶子节点路径 4->9->5 代表数字 495
//从根到叶子节点路径 4->9->1 代表数字 491
//从根到叶子节点路径 4->0 代表数字 40
//因此，数字总和 = 495 + 491 + 40 = 1026
//
//
//提示：
//
//树中节点的数目在范围 [1, 1000] 内
//0 <= Node.val <= 9
//树的深度不超过 10
//
//
//注意：本题与主站 129 题相同： https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
//
//通过次数5,573提交次数7,394
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/2/22 9:52 PM
package _Etpl5

import . "leetcode-go/container/tree"

func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, num int) int {
		if root == nil {
			return 0
		}
		num = num*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return num
		}
		return dfs(root.Left, num) + dfs(root.Right, num)
	}
	return dfs(root, 0)
}
