package topic0297

import (
	"strconv"
	"strings"
)

/*
297. 二叉树的序列化与反序列化
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。



示例 1：


输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：

输入：root = [1,2]
输出：[1,2]


提示：

树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000
通过次数98,039提交次数176,142

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/20/21 4:18 PM
*/

/**
 * Definition for a binary container node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import . "leetcode-go/container/tree"

const (
	Null = "X"
	Sep  = ","
)

//// 广度遍历解法
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	builder := strings.Builder{}
//	q := []*TreeNode{root}
//	for len(q) > 0 {
//		if builder.Len() > 0 {
//			builder.WriteString(Sep)
//		}
//		node := q[0]
//		q = q[1:]
//		c := Null
//		if node != nil {
//			c = strconv.Itoa(node.Val)
//		}
//		builder.WriteString(c)
//		if node != nil {
//			q = append(q, node.Left)
//			q = append(q, node.Right)
//		}
//	}
//	return builder.String()
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	if data == "X" {
//		return nil
//	}
//	nodes := strings.Split(data, Sep)
//	root := parseNode(nodes[0])
//	q := []*TreeNode{root}
//	i := 1
//	for i < len(nodes) {
//		node := q[0]
//		q = q[1:]
//		node.Left = parseNode(nodes[i])
//		node.Right = parseNode(nodes[i+1])
//		i += 2
//		if node.Left != nil {
//			q = append(q, node.Left)
//		}
//		if node.Right != nil {
//			q = append(q, node.Right)
//		}
//	}
//	return root
//}
//
//func parseNode(s string) *TreeNode {
//	if s == "X" {
//		return nil
//	}
//	i, _ := strconv.Atoi(s)
//	return &TreeNode{Val: i}
//}

//// 深度遍历解法
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	if root == nil {
//		return "X"
//	}
//	builder := strings.Builder{}
//	var stack []*TreeNode
//	for root != nil || len(stack) > 0 {
//		if root != nil {
//			if builder.Len() > 0 {
//				builder.WriteString(Sep)
//			}
//			builder.WriteString(strconv.Itoa(root.Val))
//			stack = append(stack, root)
//			root = root.Left
//		} else {
//			builder.WriteString(Sep)
//			builder.WriteString(Null)
//			root = stack[len(stack)-1].Right
//			stack = stack[:len(stack)-1]
//		}
//	}
//	return builder.String()
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	if data == "X" {
//		return nil
//	}
//	nodes := strings.Split(data, Sep)
//	root := parseNode(nodes[0])
//	stack := []*TreeNode{root}
//	i := 1
//	left := true
//	for i < len(nodes) {
//		node := parseNode(nodes[i])
//		i++
//		if left {
//			stack[len(stack)-1].Left = node
//		} else {
//			stack[len(stack)-1].Right = node
//			stack = stack[:len(stack)-1]
//		}
//		left = node != nil
//		if node != nil {
//			stack = append(stack, node)
//		}
//	}
//	return root
//}
//
//func parseNode(s string) *TreeNode {
//	if s == "X" {
//		return nil
//	}
//	i, _ := strconv.Atoi(s)
//	return &TreeNode{Val: i}
//}

// 括号表示编码 + 递归下降解码,参考了官方题解
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("X")
			return
		}
		sb.WriteString("(")
		dfs(root.Left)
		sb.WriteString(")")
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteString("(")
		dfs(root.Right)
		sb.WriteString(")")
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	i := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if data[i] == 'X' {
			i++
			return nil
		}
		root := &TreeNode{}
		i++
		root.Left = dfs()
		i++
		j := i + 1
		for data[j] != '(' {
			j++
		}
		num, _ := strconv.Atoi(data[i:j])
		root.Val = num
		i = j + 1
		root.Right = dfs()
		i++
		return root
	}
	return dfs()
}
