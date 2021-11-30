// 面试题 04.09. 二叉搜索树序列
//从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。给定一个由不同节点组成的二叉搜索树，输出所有可能生成此树的数组。
//
//
//
//示例：
//给定如下二叉树
//
//        2
//       / \
//      1   3
//返回：
//
//[
//   [2,1,3],
//   [2,3,1]
//]
//通过次数6,987提交次数14,649
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/30/21 2:33 PM
package bst_sequences_lcci

import (
	"container/list"
	"leetcode-go/tree"
)

// 过于复杂
func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	q := []*TreeNode{root}
	ancestors := map[*TreeNode]*TreeNode{}
	childCount := map[*TreeNode]int{}
	var leaves []*TreeNode
	count := 0
	for len(q) > 0 {
		size := len(q)
		count += size
		for i := 0; i < size; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
				ancestors[node.Left] = node
				childCount[node]++
			}
			if node.Right != nil {
				q = append(q, node.Right)
				ancestors[node.Right] = node
				childCount[node]++
			}
			if node.Left == nil && node.Right == nil {
				leaves = append(leaves, node)
			}
		}
		q = q[size:]
	}
	var arrays [][]int
	backtrack(leaves, ancestors, childCount, &arrays, make([]int, 0, count))
	return arrays
}

func backtrack(leaves []*TreeNode, ancestors map[*TreeNode]*TreeNode, childCount map[*TreeNode]int, arrays *[][]int, arr []int) {
	if len(leaves) == 0 {
		a := make([]int, len(arr))
		copy(a, arr)
		l, r := 0, len(a)-1
		for l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
		*arrays = append(*arrays, a)
		return
	}

	for i, leaf := range leaves {
		newLeaves := make([]*TreeNode, len(leaves))
		copy(newLeaves, leaves)
		newLeaves = append(newLeaves[:i], newLeaves[i+1:]...)
		ancestor := ancestors[leaf]
		childCount[ancestor]--
		if childCount[ancestor] == 0 {
			newLeaves = append(newLeaves, ancestor)
		}
		backtrack(newLeaves, ancestors, childCount, arrays, append(arr, leaf.Val))
		childCount[ancestor]++
	}
}

func BSTSequences2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	var ans [][]int
	var path []int
	candidates := []*TreeNode{root}
	var dfs func()
	dfs = func() {
		if len(candidates) == 0 {
			p := make([]int, len(path))
			copy(p, path)
			ans = append(ans, p)
			return
		}
		length := len(candidates)
		for i := 0; i < length; i++ {
			node := candidates[0]
			candidates = candidates[1:]
			if node.Left != nil {
				candidates = append(candidates, node.Left)
			}
			if node.Right != nil {
				candidates = append(candidates, node.Right)
			}
			path = append(path, node.Val)
			dfs()
			candidates = candidates[:length-1]
			candidates = append(candidates, node)
			path = path[:len(path)-1]
		}
	}
	dfs()
	return ans
}

// 思路和解法2一样，区别在于这里使用链表
func BSTSequences3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	var ans [][]int
	var path []int
	candidates := list.New()
	candidates.PushBack(root)
	var dfs func()
	dfs = func() {
		if candidates.Len() == 0 {
			p := make([]int, len(path))
			copy(p, path)
			ans = append(ans, p)
			return
		}
		length := candidates.Len()
		for i := 0; i < length; i++ {
			node := candidates.Remove(candidates.Front()).(*TreeNode)
			if node.Left != nil {
				candidates.PushBack(node.Left)
			}
			if node.Right != nil {
				candidates.PushBack(node.Right)
			}
			path = append(path, node.Val)
			dfs()
			for candidates.Len() > length-1 {
				candidates.Remove(candidates.Back())
			}
			candidates.PushBack(node)
			path = path[:len(path)-1]
		}
	}
	dfs()
	return ans
}
