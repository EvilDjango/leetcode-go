// 剑指 Offer II 050. 向下的路径节点之和
//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
//路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//
//
//示例 1：
//
//
//
//输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//输出：3
//解释：和等于 8 的路径有 3 条，如图所示。
//示例 2：
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：3
//
//
//提示:
//
//二叉树的节点个数的范围是 [0,1000]
//-109 <= Node.val <= 109
//-1000 <= targetSum <= 1000
//
//
//注意：本题与主站 437 题相同：https://leetcode-cn.com/problems/path-sum-iii/
//
//通过次数4,440提交次数7,390
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/2/22 9:22 PM
package _eUYwP

import . "leetcode-go/container/tree"

func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, paths []int) {
		if node == nil {
			return
		}
		if node.Val == targetSum {
			ans++
		}
		newPaths := make([]int, len(paths)+1)
		for i, path := range paths {
			newPaths[i] = path + node.Val
			if newPaths[i] == targetSum {
				ans++
			}
		}
		newPaths[len(newPaths)-1] = node.Val
		dfs(node.Left, newPaths)
		dfs(node.Right, newPaths)
	}
	dfs(root, nil)
	return ans
}

// 只记录前缀和
func pathSum2(root *TreeNode, targetSum int) int {
	ans := 0
	preSum := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		ans += preSum[sum-targetSum]
		preSum[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		preSum[sum]--
	}
	dfs(root, 0)
	return ans
}
