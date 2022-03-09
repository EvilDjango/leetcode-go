// 面试题32 - I. 从上到下打印二叉树
//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//
//
//
//例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回：
//
//[3,9,20,15,7]
//
//
//提示：
//
//节点总数 <= 1000
//通过次数179,340提交次数280,561
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/9 上午10:37
package cong_shang_dao_xia_da_yin_er_cha_shu_lcof

import . "leetcode-go/container/tree"

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var ans []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		ans = append(ans, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return ans
}
