// 剑指 Offer II 013. 二维子矩阵的和
//给定一个二维矩阵 matrix，以下类型的多个请求：
//
//计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
//实现 NumMatrix 类：
//
//NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
//int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。
//
//
//示例 1：
//
//
//
//输入:
//["NumMatrix","sumRegion","sumRegion","sumRegion"]
//[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
//输出:
//[null, 8, 11, 12]
//
//解释:
//NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]]);
//numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
//numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
//numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
//
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 200
//-105 <= matrix[i][j] <= 105
//0 <= row1 <= row2 < m
//0 <= col1 <= col2 < n
//最多调用 104 次 sumRegion 方法
//
//
//注意：本题与主站 304 题相同： https://leetcode-cn.com/problems/range-sum-query-2d-immutable/
//
//通过次数11,940提交次数18,003
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/22 上午11:51
package O4NDxx

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		sums[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := matrix[i][j]
			sum += sums[i][j+1]
			sum += sums[i+1][j]
			sum -= sums[i][j]
			sums[i+1][j+1] = sum
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] + this.sums[row1][col1] - this.sums[row1][col2+1] - this.sums[row2+1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
