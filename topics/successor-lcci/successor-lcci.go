// 面试题 04.06. 后继者
//设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
//
//如果指定节点没有对应的“下一个”节点，则返回null。
//
//示例 1:
//
//输入: root = [2,1,3], p = 1
//
//  2
// / \
//1   3
//
//输出: 2
//示例 2:
//
//输入: root = [5,3,6,2,4,null,null,1], p = 6
//
//      5
//     / \
//    3   6
//   / \
//  2   4
// /
//1
//
//输出: null
//通过次数21,398提交次数36,088
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/1/21 12:55 PM
package successor_lcci

import "leetcode-go/container"

// 把root当成一个普通二叉树来进行中序遍历，循环
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var q []*TreeNode
	var lastVisit *TreeNode
	for len(q) > 0 || root != nil {
		if root != nil {
			q = append(q, root)
			root = root.Left
		} else {
			root = q[len(q)-1]
			q = q[:len(q)-1]
			if lastVisit == p {
				return root
			}
			lastVisit = root
			root = root.Right
		}
	}
	return nil
}

// 把root当成一个普通二叉树来进行中序遍历，递归
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var lastVisit, successor *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || successor != nil {
			return
		}
		dfs(root.Left)
		if lastVisit == p {
			successor = root
		}
		lastVisit = root
		dfs(root.Right)
	}
	dfs(root)
	return successor
}

// 把root当成一个普通二叉树来进行莫里斯中序遍历
func inorderSuccessor3(root *TreeNode, p *TreeNode) *TreeNode {
	var lastVisit *TreeNode
	for root != nil {
		if root.Left == nil {
			if lastVisit == p {
				return root
			}
			lastVisit = root
			root = root.Right
		} else {
			rightMost := root.Left
			for rightMost.Right != nil && rightMost.Right != root {
				rightMost = rightMost.Right
			}
			if rightMost.Right == nil {
				// 搭桥，左子树遍历完后由此回到根节点
				rightMost.Right = root
				root = root.Left
			} else {
				// 发现之前搭的桥，我们再次回到根节点说明左子树已经遍历完成，拆桥复原
				rightMost.Right = nil
				if lastVisit == p {
					return root
				}
				lastVisit = root
				root = root.Right
			}
		}
	}
	return nil
}

// 利用二叉搜索树的性质来快速查找,循环
func inorderSuccessor4(root *TreeNode, p *TreeNode) *TreeNode {
	var option *TreeNode
	for root != nil {
		if root.Val > p.Val {
			option = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return option
}

// 利用二叉搜索树的性质来快速查找，递归
func inorderSuccessor5(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	} else {
		res := inorderSuccessor(root.Left, p)
		if res != nil {
			return res
		}
		return root
	}
}
