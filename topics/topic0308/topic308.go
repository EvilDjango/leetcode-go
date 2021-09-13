// 308. 二维区域和检索 - 可变
//给你一个二维矩阵 matrix ，你需要处理下面两种类型的若干次查询：
//
//更新：更新 matrix 中某个单元的值。
//求和：计算矩阵 matrix 中某一矩形区域元素的 和 ，该区域由 左上角 (row1, col1) 和 右下角 (row2, col2) 界定。
//实现 NumMatrix 类：
//
//NumMatrix(int[][] matrix) 用整数矩阵 matrix 初始化对象。
//void update(int row, int col, int val) 更新 matrix[row][col] 的值到 val 。
//int sumRegion(int row1, int col1, int row2, int col2) 返回矩阵 matrix 中指定矩形区域元素的 和 ，该区域由 左上角 (row1, col1) 和 右下角 (row2, col2) 界定。
//
//
//示例：
//
//
//输入
//["NumMatrix", "sumRegion", "update", "sumRegion"]
//[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [3, 2, 2], [2, 1, 4, 3]]
//输出
//[null, 8, null, 10]
//
//解释
//NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
//numMatrix.sumRegion(2, 1, 4, 3); // 返回 8 (即, 左侧红色矩形的和)
//numMatrix.update(3, 2, 2);       // 矩阵从左图变为右图
//numMatrix.sumRegion(2, 1, 4, 3); // 返回 10 (即，右侧红色矩形的和)
//
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 200
//-105 <= matrix[i][j] <= 105
//0 <= row < m
//0 <= col < n
//-105 <= val <= 105
//0 <= row1 <= row2 < m
//0 <= col1 <= col2 < n
//最多调用104 次 sumRegion 和 update 方法
//通过次数2,146提交次数3,544
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/7/21 8:47 PM
package topic0308

// 树状数组
type NumMatrix struct {
	nums, tree []int
	rows, cols int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	N := m * n
	nums, tree := make([]int, N), make([]int, N+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j
			nums[index] = matrix[i][j]
			tree[index+1] += nums[index]
			index2 := index + 1 + lowBit(index+1)
			if index2 < len(tree) {
				tree[index2] += tree[index+1]
			}
		}
	}
	return NumMatrix{nums, tree, m, n}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	index := row*this.cols + col
	delta := val - this.nums[index]
	for i := index + 1; i < len(this.tree); i += lowBit(i) {
		this.tree[i] += delta
	}
	this.nums[index] = val
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.sumRegion(i*this.cols+col1, i*this.cols+col2)
	}
	return sum
}

func (this *NumMatrix) sumRegion(i, j int) int {
	return this.sum(j) - this.sum(i-1)
}

func (this *NumMatrix) sum(i int) int {
	sum := 0
	for i := i + 1; i > 0; i -= lowBit(i) {
		sum += this.tree[i]
	}
	return sum
}

func lowBit(i int) int {
	return i & -i
}
