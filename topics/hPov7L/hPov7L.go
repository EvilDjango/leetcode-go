// 剑指 Offer II 044. 二叉树每层的最大值
//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
//
//
//
//示例1：
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//解释:
//          1
//         / \
//        3   2
//       / \   \
//      5   3   9
//示例2：
//
//输入: root = [1,2,3]
//输出: [1,3]
//解释:
//          1
//         / \
//        2   3
//示例3：
//
//输入: root = [1]
//输出: [1]
//示例4：
//
//输入: root = [1,null,2]
//输出: [1,2]
//解释:
//           1
//            \
//             2
//示例5：
//
//输入: root = []
//输出: []
//
//
//提示：
//
//二叉树的节点个数的范围是 [0,104]
//-231 <= Node.val <= 231 - 1
//
//
//注意：本题与主站 515 题相同： https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
//
//通过次数6,014提交次数9,402
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/3/22 2:24 PM
package hPov7L

import . "leetcode-go/container/tree"

func largestValues(root *TreeNode) []int {
	var ans []int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth+1 > len(ans) {
			ans = append(ans, root.Val)
		} else if ans[depth] < root.Val {
			ans[depth] = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
