// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/10/21 3:03 PM
package sorted_matrix_search_lcci

// 二分查找
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 || matrix[i][0] > target {
			return false
		}
		l, r := 0, len(matrix[0])
		for l < r {
			mid := (r-l)/2 + l
			if matrix[i][mid] > target {
				r = mid
			} else if matrix[i][mid] < target {
				l = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// 从左下角开始查找，那么可以把矩阵看成一个二叉搜索树，向上查找等同于查找二超搜索树的左子树，向右等同于查找右子树。
//而向下和向左是不可取的，因为下侧的数值太大（正是因为下方的数值太大，才会向上走到当前位置），而左侧的数值太小（正是因为左侧数值太小，才会向右走到当前位置）。
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	row, col := m-1, 0
	for row >= 0 && col < n {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}
