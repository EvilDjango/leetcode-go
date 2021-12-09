// 面试题 01.08. 零矩阵
//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
//
//
//示例 1：
//
//输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//示例 2：
//
//输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//通过次数47,330提交次数76,029
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 1:49 PM
package zero_matrix_lcci

// 记录需要置零的位置
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(matrix), len(matrix[0])
	toChange := map[[2]int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				continue
			}
			for _, direct := range directions {
				for x, y := i+direct[0], j+direct[1]; x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] != 0; x, y = x+direct[0], y+direct[1] {
					toChange[[2]int{x, y}] = true
				}
			}
		}
	}
	for indexes := range toChange {
		matrix[indexes[0]][indexes[1]] = 0
	}
}

// 仅记录要置零的行号和列号
func setZeroes2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 使用第一行和第一列来记录需要置零的列和行
func setZeroes3(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if row[j] == 0 {
				row[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
