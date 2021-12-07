package topic0297_03

import (
	"leetcode-go/container/tree"
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
 * type container.TreeNode struct {
 *     Val int
 *     Left *container.TreeNode
 *     Right *container.TreeNode
 * }
 */

const (
	Null = 'X'
)

type Codec struct {
}

// 括号表示编码
func Constructor() Codec {
	return Codec{}
}

// Serializes a container to a single string.
func (this *Codec) serialize(root *tree.TreeNode) string {
	if root == nil {
		return string(Null)
	}

	builder := strings.Builder{}
	builder.WriteString("(")
	builder.WriteString(this.serialize(root.Left))
	builder.WriteString(")")
	builder.WriteString(strconv.Itoa(root.Val))
	builder.WriteString("(")
	builder.WriteString(this.serialize(root.Right))
	builder.WriteString("(")
	return builder.String()
}

// Deserializes your encoded data to container.
func (this *Codec) deserialize(data string) *tree.TreeNode {
	var parse func() *tree.TreeNode
	i := 0
	parse = func() *tree.TreeNode {
		if data[i] == Null {
			i++
			return nil
		}
		node := &tree.TreeNode{}
		i++
		node.Left = parse()
		i++
		j := i
		for ; j < len(data) && data[j] != '('; j++ {
		}
		node.Val, _ = strconv.Atoi(data[i:j])
		i = j + 1
		node.Right = parse()
		i++
		return node
	}
	return parse()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
