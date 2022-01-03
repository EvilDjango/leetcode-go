// 剑指 Offer II 046. 二叉树的右侧视图
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
//
//
//示例 1:
//
//
//
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]
//示例 2:
//
//输入: [1,null,3]
//输出: [1,3]
//示例 3:
//
//输入: []
//输出: []
//
//
//提示:
//
//二叉树的节点个数的范围是 [0,100]
//-100 <= Node.val <= 100
//
//
//注意：本题与主站 199 题相同：https://leetcode-cn.com/problems/binary-tree-right-side-view/
//
//通过次数5,409提交次数7,507
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/3/22 1:52 PM
package WNC0Lk

import . "leetcode-go/container/tree"

func rightSideView(root *TreeNode) []int {
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth+1 > len(ans) {
			ans = append(ans, root.Val)
		} else {
			ans[depth] = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
