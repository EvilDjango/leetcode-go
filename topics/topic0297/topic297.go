package topic0297

import (
	"bytes"
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

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *tree.TreeNode) string {
	buf := bytes.Buffer{}
	buf.WriteString(toString(root))
	buf.WriteString(Sep)
	var q []*tree.TreeNode
	if root != nil {
		q = append(q, root)
	}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			buf.WriteString(toString(node.Left))
			buf.WriteString(Sep)
			buf.WriteString(toString(node.Right))
			buf.WriteString(Sep)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *tree.TreeNode {
	vals := strings.Split(data, Sep)
	root := parseNode(vals[0])
	i := 1
	var q []*tree.TreeNode
	if root != nil {
		q = append(q, root)
	}
	for i < len(vals) && len(q) > 0 {
		size := len(q)
		for j := 0; i < len(vals) && j < size; j++ {
			node := q[j]
			node.Left = parseNode(vals[i])
			if i+1 < len(vals) {
				node.Right = parseNode(vals[i+1])
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			i += 2
		}
		q = q[size:]
	}
	if i != len(vals) {
		panic("invalid data: " + data)
	}
	return root
}

func parseNode(s string) *tree.TreeNode {
	if s == Null {
		return nil
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid data. want digit, got " + s)
	}
	return &tree.TreeNode{Val: val}
}

func toString(node *tree.TreeNode) string {
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
