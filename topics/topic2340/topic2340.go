// 面试题 17.24. 最大子矩阵
//给定一个正整数、负整数和 0 组成的 N × M 矩阵，编写代码找出元素总和最大的子矩阵。
//
//返回一个数组 [r1, c1, r2, c2]，其中 r1, c1 分别代表子矩阵左上角的行号和列号，r2, c2 分别代表右下角的行号和列号。若有多个满足条件的子矩阵，返回任意一个均可。
//
//注意：本题相对书上原题稍作改动
//
//示例：
//
//输入：
//[
//   [-1,0],
//   [0,-1]
//]
//输出：[0,1,0,1]
//解释：输入中标粗的元素即为输出所表示的矩阵
//
//
//说明：
//
//1 <= matrix.length, matrix[0].length <= 200
//通过次数9,782提交次数18,845
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/10/21 4:06 PM
package topic2340

import "math"

func maxSubArray(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	maxSum := math.MinInt32
	var ans []int
	for i := 0; i < n; i++ {
		col := make([]int, m)
		for j := i; j < n; j++ {
			sum := 0
			startRow := 0
			for k := 0; k < m; k++ {
				col[k] += matrix[k][j]
				if sum < 0 {
					sum = 0
					startRow = k
				}
				sum += col[k]
				if sum > maxSum {
					maxSum = sum
					ans = []int{startRow, i, k, j}
				}
			}
		}
	}
	return ans
}
