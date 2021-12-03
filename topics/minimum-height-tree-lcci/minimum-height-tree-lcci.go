// 面试题 04.02. 最小高度树
//给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
//
//示例:
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//          0
//         / \
//       -3   9
//       /   /
//     -10  5
//通过次数36,053提交次数45,669
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/3/21 2:20 PM
package minimum_height_tree_lcci

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	pivot := len(nums) / 2
	root := &TreeNode{Val: nums[pivot]}
	root.Left = sortedArrayToBST(nums[:pivot])
	root.Right = sortedArrayToBST(nums[pivot+1:])
	return root
}
