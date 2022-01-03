// 剑指 Offer II 056. 二叉搜索树中两个节点之和
//给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。假设二叉搜索树中节点的值均唯一。
//
//
//
//示例 1：
//
//输入: root = [8,6,10,5,7,9,11], k = 12
//输出: true
//解释: 节点 5 和节点 7 之和等于 12
//示例 2：
//
//输入: root = [8,6,10,5,7,9,11], k = 22
//输出: false
//解释: 不存在两个节点值之和为 22 的节点
//
//
//提示：
//
//二叉树的节点个数的范围是  [1, 104].
//-104 <= Node.val <= 104
//root 为二叉搜索树
//-105 <= k <= 105
//
//
//注意：本题与主站 653 题相同： https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
//
//通过次数6,065提交次数8,412
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/31/21 9:32 PM
package opLdQZ

import . "leetcode-go/container/tree"

// 利用哈希表，一次遍历
func findTarget(root *TreeNode, k int) bool {
	seen := map[int]bool{}
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if seen[k-root.Val] {
			return true
		}
		seen[root.Val] = true
		return dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}

// 利用二叉搜索树的性质
func findTarget2(root *TreeNode, k int) bool {
	var find func(*TreeNode, int) bool
	find = func(root *TreeNode, target int) bool {
		if root == nil {
			return false
		}
		if root.Val < target {
			return find(root.Right, target)
		} else if root.Val > target {
			return find(root.Left, target)
		} else {
			return true
		}
	}
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}
		rest := target - node.Val
		return (rest != node.Val && find(root, rest)) || dfs(node.Left, target) || dfs(node.Right, target)
	}
	return dfs(root, k)
}

type InorderIterator struct {
	curr  *TreeNode
	stack []*TreeNode
}

func (it *InorderIterator) Next() *TreeNode {
	curr := it.curr
	stack := it.stack
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}
	node := stack[len(stack)-1]
	it.stack = stack[:len(stack)-1]
	it.curr = node.Right
	return node
}

type ReverseInorderIterator struct {
	curr  *TreeNode
	stack []*TreeNode
}

func (it *ReverseInorderIterator) Next() *TreeNode {
	curr := it.curr
	stack := it.stack
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Right
	}
	node := stack[len(stack)-1]
	it.stack = stack[:len(stack)-1]
	it.curr = node.Left
	return node
}

// 双指针
func findTarget3(root *TreeNode, k int) bool {
	it, it2 := InorderIterator{curr: root}, ReverseInorderIterator{curr: root}
	left, right := it.Next(), it2.Next()
	for left.Val < right.Val {
		if left.Val+right.Val > k {
			right = it2.Next()
		} else if left.Val+right.Val < k {
			left = it.Next()
		} else {
			return true
		}
	}
	return false
}
