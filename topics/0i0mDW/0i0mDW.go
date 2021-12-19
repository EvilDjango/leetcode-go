// 剑指 Offer II 099. 最小路径之和
//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：一个机器人每次只能向下或者向右移动一步。
//
//
//
//示例 1：
//
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//示例 2：
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 200
//0 <= grid[i][j] <= 100
//
//
//注意：本题与主站 64 题相同： https://leetcode-cn.com/problems/minimum-path-sum/
//
//通过次数4,356提交次数6,118
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/19/21 1:26 PM
package _i0mDW

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			if dp[j-1] < dp[j] {
				dp[j] = dp[j-1]
			}
			dp[j] += grid[i][j]
		}
	}
	return dp[n-1]
}
