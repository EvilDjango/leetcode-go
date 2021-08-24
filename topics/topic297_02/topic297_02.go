package topic297_02

import (
	"leetcode-go/tree"
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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const (
	Sep  = "|"
	Null = ""
)

type Codec struct {
}

// 深度遍历：前序遍历
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return Null
	}
	builder := strings.Builder{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		builder.WriteString(toString(root))
		builder.WriteString(Sep)
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, Sep)
	root := parseNode(vals[0])
	stack := []*TreeNode{root}
	left := true
	for i := 1; i < len(vals); i++ {
		node := parseNode(vals[i])
		ans := stack[len(stack)-1]
		if left {
			ans.Left = node
		} else {
			ans.Right = node
			stack = stack[:len(stack)-1]
		}
		if node == nil {
			left = false
		} else {
			left = true
			stack = append(stack, node)
		}
	}
	return root
}

func parseNode(s string) *TreeNode {
	if s == Null {
		return nil
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid data. want digit, got " + s)
	}
	return &TreeNode{Val: val}
}

func toString(node *TreeNode) string {
	if node == nil {
		return Null
	}
	return strconv.Itoa(node.Val)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
