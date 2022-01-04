// 剑指 Offer II 040. 矩阵中最大的矩形
//给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
//
//注意：此题 matrix 输入格式为一维 01 字符串数组。
//
//
//
//示例 1：
//
//
//
//输入：matrix = ["10100","10111","11111","10010"]
//输出：6
//解释：最大矩形如上图所示。
//示例 2：
//
//输入：matrix = []
//输出：0
//示例 3：
//
//输入：matrix = ["0"]
//输出：0
//示例 4：
//
//输入：matrix = ["1"]
//输出：1
//示例 5：
//
//输入：matrix = ["00"]
//输出：0
//
//
//提示：
//
//rows == matrix.length
//cols == matrix[0].length
//0 <= row, cols <= 200
//matrix[i][j] 为 '0' 或 '1'
//
//
//注意：本题与主站 85 题相同（输入参数格式不同）： https://leetcode-cn.com/problems/maximal-rectangle/
//
//通过次数3,218提交次数5,599
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/4/22 1:47 PM
package PLYXKQ

// 动态规划得到每个位置向右的最长1序列，然后遍历得到柱状图中的最大矩形
func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	// dp[i][j]表示从(i,j)出发向右最长的连续1序列
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i][j+1] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			length := dp[i][j]
			for k := i; k < m && dp[k][j] != 0; k++ {
				if length > dp[k][j] {
					length = dp[k][j]
				}
				if length*(k-i+1) > ans {
					ans = length * (k - i + 1)
				}
			}
		}
	}
	return ans
}

// 动态规划得到每个位置向右的最长1序列，然后利用单调栈得到柱状图中的最大矩形
func maximalRectangle2(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	// dp[i][j]表示从(i,j)出发向下最长的连续1序列
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i+1][j] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		area := maxRectangle(dp[i])
		if area > ans {
			ans = area
		}
	}
	return ans
}

func maxRectangle(heights []int) int {
	ans := 0
	stack := []int{-1}
	for i, h := range heights {
		j := len(stack) - 1
		for ; j > 0 && heights[stack[j]] >= h; j-- {
			area := (i - stack[j-1] - 1) * heights[stack[j]]
			ans = max(ans, area)
		}
		stack = stack[:j+1]
		stack = append(stack, i)
	}
	for i := len(stack) - 1; i > 0; i-- {
		area := (len(heights) - stack[i-1] - 1) * heights[stack[i]]
		ans = max(ans, area)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
