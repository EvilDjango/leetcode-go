package tree

import (
	"bytes"
	"math"
	"math/rand"
	"strconv"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/5/21 5:30 PM
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const None = math.MaxInt64

func (t *TreeNode) String() string {
	buf := bytes.NewBufferString("\n")
	q := []*TreeNode{t}
	children := 1
	for len(q) > 0 && children > 0 {
		children = 0
		size := len(q)
		for i := 0; i < size; i++ {
			t = q[i]
			if i > 0 {
				buf.WriteString("  ")
			}
			if t != nil {
				buf.WriteString(strconv.Itoa(t.Val))
				if t.Left != nil || t.Right != nil {
					children++
				}
				q = append(q, t.Left)
				q = append(q, t.Right)
			} else {
				buf.WriteString("X")
			}
		}
		q = q[size:]
		buf.WriteString("\n")
	}
	return buf.String()
}

func (t *TreeNode) Clone() *TreeNode {
	if t == nil {
		return nil
	}
	newRoot := &TreeNode{Val: t.Val}
	newRoot.Left = t.Left.Clone()
	newRoot.Right = t.Right.Clone()
	return newRoot
}

func (t *TreeNode) Len() int {
	if t == nil {
		return 0
	}
	return t.Left.Len() + t.Right.Len() + 1
}

func (t *TreeNode) Equals(q *TreeNode) bool {
	if t == nil && q == nil {
		return true
	}
	if t == nil || q == nil {
		return t == q
	}

	if t.Val != q.Val {
		return false
	}
	return t.Left.Equals(q.Left) && t.Right.Equals(q.Right)
}

func (t *TreeNode) GetChild(val int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == val {
		return t
	}
	child := t.Left.GetChild(val)
	if child != nil {
		return child
	}
	return t.Right.GetChild(val)
}

func New(nums ...int) *TreeNode {
	n := len(nums)
	if n == 0 || nums[0] == None {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 {
		size := len(q)
		for j := 0; j < size && i < n; j++ {
			cur := q[j]
			if nums[i] != None {
				cur.Left = &TreeNode{Val: nums[i]}
				q = append(q, cur.Left)
			}
			if i+1 < n && nums[i+1] != None {
				cur.Right = &TreeNode{Val: nums[i+1]}
				q = append(q, cur.Right)
			}
			i += 2
		}
		q = q[size:]
	}
	return root
}

// 随机生成一棵树，节点数范围是[1,n]
func Random(n int) (*TreeNode, int) {
	root := &TreeNode{Val: 0}
	q := []*TreeNode{root}
	index := 1
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size && index < n; i++ {
			cur := q[i]
			if rand.Intn(10) >= 5 {
				cur.Left = &TreeNode{Val: index}
				q = append(q, cur.Left)
				index++
			}
			if rand.Intn(10) >= 5 {
				cur.Right = &TreeNode{Val: index}
				q = append(q, cur.Right)
				index++
			}
		}
		q = q[size:]
	}
	return root, index
}
