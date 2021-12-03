// 面试题 04.08. 首个共同祖先
//设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。
//
//例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
//
//    3
//   / \
//  5   1
// / \ / \
//6  2 0  8
//  / \
// 7   4
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
//说明:
//
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉树中。
//通过次数17,718提交次数24,716
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/1/21 9:58 AM
package first_common_ancestor_lcci

import "leetcode-go/tree"

func lowestCommonAncestor(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode {
	path1 := getPath(root, p)
	path2 := getPath(root, q)
	minLen := len(path1)
	if len(path2) < minLen {
		minLen = len(path2)
	}
	i := 0
	for ; i < minLen && path1[i] == path2[i]; i++ {
		// empty body
	}
	return path1[i-1]
}

func getPath(root, target *tree.TreeNode) []*tree.TreeNode {
	if root == nil {
		return nil
	}
	path := []*tree.TreeNode{root}
	if root == target {
		return path
	}
	subPath := getPath(root.Left, target)
	if subPath == nil {
		subPath = getPath(root.Right, target)
	}
	if subPath != nil {
		return append(path, subPath...)
	}
	return nil
}

func lowestCommonAncestor2(root *tree.TreeNode, p *tree.TreeNode, q *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
