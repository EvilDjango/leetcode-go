package topics

import (
	"reflect"
	"strconv"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/5/21 7:29 PM
*/

func TestNew(t *testing.T) {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Left.Right = &TreeNode{Val: 7}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"",
			args{[]int{0, 1, 2, 3, None, 4, None, None, 5, 6, 7}},
			root,
		},
		{
			"",
			args{[]int{}},
			nil,
		},
		{
			"",
			args{[]int{None}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	root, _ := Random(100)
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{nil, nil},
			true,
		},
		{
			"",
			args{nil, &TreeNode{Val: 1}},
			false,
		},
		{
			"",
			args{&TreeNode{Val: 1}, &TreeNode{Val: 2}},
			false,
		},
		{
			"",
			args{root, root},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.p.Equals(tt.args.q); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			root, size := Random(100)
			realSize := root.Len()
			if realSize != size {
				t.Errorf("Random tree size = %v, want %v", realSize, size)
			}
		})
	}
}

func Test_Len(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{root},
			2,
		},
		{
			"",
			args{root.Left},
			1,
		},
		{
			"",
			args{nil},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.root.Len(); got != tt.want {
				t.Errorf("size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t1, _ := Random(100)
			t2 := t1.Clone()
			if !reflect.DeepEqual(t1, t2) {
				t.Errorf("size() = %v, want %v", t1, t2)
			}
		})
	}
}

func TestTreeNode_String(t1 *testing.T) {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	tests := []struct {
		name string
		t    *TreeNode
		want string
	}{
		{
			"",
			nil,
			"\nX\n",
		},
		{
			"",
			root,
			"\n0\n1  2\n",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
