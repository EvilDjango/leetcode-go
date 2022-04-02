// 剑指 Offer 68 - II. 二叉树的最近公共祖先
//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
//
//
//
//
//
//示例 1:
//
//输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出: 3
//解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//示例 2:
//
//输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//
//
//说明:
//
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉树中。
//注意：本题与主站 236 题相同：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
//
//通过次数153,142提交次数217,898
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/4/2 下午4:48
package er_cha_shu_de_zui_jin_gong_gong_zu_xian_lcof

import (
	. "leetcode-go/container/tree"
)

// 自己写的递归算法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	// 返回值表示root下包含的p或q的总数量
	var dfs func(root, p, q *TreeNode) int
	dfs = func(root, p, q *TreeNode) int {
		if root == nil {
			return 0
		}
		cnt := dfs(root.Left, p, q) + dfs(root.Right, p, q)
		if root == p || root == q {
			cnt++
		}
		if ans == nil && cnt == 2 {
			ans = root
		}
		return cnt
	}
	dfs(root, p, q)
	return ans
}

// 参考官方题解写的递归算法
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
