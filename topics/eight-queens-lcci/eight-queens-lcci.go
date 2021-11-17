// 面试题 08.12. 八皇后
//设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。
//
//注意：本题相对原题做了扩展
//
//示例:
//
// 输入：4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//通过次数17,453提交次数22,939
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/17/21 2:58 PM
package eight_queens_lcci

import (
	"math/bits"
)

// 回溯 比较笨的办法，空间占用很高
func solveNQueens(n int) [][]string {
	var ans [][]string
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
	}
	backtrack(grid, &ans, 0)
	return ans
}

func backtrack(grid [][]byte, ans *[][]string, i int) {
	n := len(grid)
	if i == n {
		result := make([]string, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] != 'Q' {
					grid[i][j] = '.'
				}
			}
			result[i] = string(grid[i])
		}
		*ans = append(*ans, result)
		return
	}
	for j := 0; j < n; j++ {
		if grid[i][j] == 0 {
			g := copyGrid(grid)
			g[i][j] = 'Q'
			markForbiddenAreaBelow(g, i, j)
			backtrack(g, ans, i+1)
		}
	}
}

// 标记下方不能放皇后的格子
func markForbiddenAreaBelow(grid [][]byte, row, col int) {
	n := len(grid)
	for i := row + 1; i < n; i++ {
		grid[i][col] = 'X'
	}
	// 左侧对角线
	for i, j := row+1, col-1; i < n && j >= 0; i, j = i+1, j-1 {
		grid[i][j] = 'X'
	}
	// 右侧对角线
	for i, j := row+1, col+1; i < n && j < n; i, j = i+1, j+1 {
		grid[i][j] = 'X'
	}
}

func copyGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

// 优化的回溯解法 依次确定每一行的皇后位置，使用位运算来判断某个位置的所在的列和对角线是否已经有一个皇后了
func solveNQueens2(n int) [][]string {
	var ans [][]string
	// queenCols[i]表示i行的皇后所在的列
	queenCols := make([]int, n)
	// 使用三个二进制数来表示列，左上-右下对角线，左下-右上对角线的占用情况，某一位为0表示未占用
	var usedCols, usedDiagonals1, usedDiagonals2 int64
	dfs(queenCols, &ans, usedCols, usedDiagonals1, usedDiagonals2, 0)
	return ans
}

func dfs(queenCols []int, ans *[][]string, usedCols int64, usedDiagonals1 int64, usedDiagonals2 int64, row int) {
	n := len(queenCols)
	if row == n {
		*ans = append(*ans, buildBoard(queenCols))
		return
	}
	for col := 0; col < n; col++ {
		diagonal1 := row - col + n
		diagonal2 := row + col
		one := int64(1)
		mask1, mask2, mask3 := one<<col, one<<diagonal1, one<<diagonal2
		if usedCols&mask1+usedDiagonals1&mask2+usedDiagonals2&mask3 == 0 {
			queenCols[row] = col
			dfs(queenCols, ans, usedCols|mask1, usedDiagonals1|mask2, usedDiagonals2|mask3, row+1)
		}
	}
}

func buildBoard(queenCols []int) []string {
	n := len(queenCols)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		bytes := make([]byte, n)
		for j := 0; j < n; j++ {
			if queenCols[i] == j {
				bytes[j] = 'Q'
			} else {
				bytes[j] = '.'
			}
		}
		result[i] = string(bytes)
	}
	return result
}

// 优化的回溯解法 还是使用位运算，但是更巧妙一些,思路参见官方题解
func solveNQueens3(n int) [][]string {
	var ans [][]string
	// queenCols[i]表示i行的皇后所在的列
	queenCols := make([]int, n)
	dfs2(queenCols, &ans, 0, 0, 0, 0)
	return ans
}

func dfs2(queenCols []int, ans *[][]string, column, diagonal1, diagonal2 int, row int) {
	n := len(queenCols)
	if row == n {
		*ans = append(*ans, buildBoard(queenCols))
		return
	}
	// 可用的列，现在1表示空闲的列
	availablePositions := (1<<n - 1) &^ (column | diagonal1 | diagonal2)
	for availablePositions > 0 {
		position := availablePositions & -availablePositions
		availablePositions &= availablePositions - 1
		col := bits.OnesCount(uint(position - 1))
		queenCols[row] = col
		dfs2(queenCols, ans, column|position, (diagonal1|position)>>1, (diagonal2|position)<<1, row+1)
	}
}
