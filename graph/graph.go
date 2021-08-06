package graph

/*
图节点

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 2:49 PM
*/

type Node struct {
	Val      int
	Adjacent []*Node
}
