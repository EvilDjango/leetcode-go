// 面试题 04.12. 求和路径
//给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//返回:
//
//3
//解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
//提示：
//
//节点总数 <= 10000
//通过次数18,582提交次数38,213
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/30/21 1:28 PM
package paths_with_sum_lcci

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	ans := 0
	find(root, sum, nil, &ans)
	return ans
}

func find(curr *TreeNode, sum int, sums []int, count *int) {
	if curr == nil {
		return
	}
	newSums := make([]int, len(sums)+1)
	copy(newSums, sums)
	for i := range newSums {
		newSums[i] += curr.Val
		if newSums[i] == sum {
			*count++
		}
	}
	find(curr.Left, sum, newSums, count)
	find(curr.Right, sum, newSums, count)
}

// 上面的解法中记录了全部子路径的和，有点浪费空间，实际上当我们遍历某一条路径时，只需要记录每个节点到根节点的路径和即可，两个节点的根路径和只差即为子路径的和。
func pathSum2(root *TreeNode, sum int) int {
	ans := 0
	dfs(root, sum, map[int]int{0: 1}, &ans, 0)
	return ans
}

func dfs(node *TreeNode, target int, sumCount map[int]int, count *int, currSum int) {
	if node == nil {
		return
	}
	currSum += node.Val
	*count += sumCount[currSum-target]
	sumCount[currSum]++
	dfs(node.Left, target, sumCount, count, currSum)
	dfs(node.Right, target, sumCount, count, currSum)
	sumCount[currSum]++
}
