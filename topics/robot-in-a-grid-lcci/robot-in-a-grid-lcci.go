// 面试题 08.02. 迷路的机器人
//设想有个机器人坐在一个网格的左上角，网格 r 行 c 列。机器人只能向下或向右移动，但不能走到一些被禁止的网格（有障碍物）。设计一种算法，寻找机器人从左上角移动到右下角的路径。
//
//
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//返回一条可行的路径，路径由经过的网格的行号和列号组成。左上角为 0 行 0 列。如果没有可行的路径，返回空数组。
//
//示例 1:
//
//输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: [[0,0],[0,1],[0,2],[1,2],[2,2]]
//解释:
//输入中标粗的位置即为输出表示的路径，即
//0行0列（左上角） -> 0行1列 -> 0行2列 -> 1行2列 -> 2行2列（右下角）
//说明：r 和 c 的值均不超过 100。
//
//通过次数13,010提交次数36,055
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/28/21 3:23 PM
package robot_in_a_grid_lcci

// 朴素的回溯
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	return backtrack(obstacleGrid, 0, 0)
}

func backtrack(grid [][]int, row, col int) [][]int {
	m, n := len(grid), len(grid[0])
	if m == row || n == col || grid[row][col] == 1 {
		return nil
	}
	path := [][]int{{row, col}}
	// 防止重复遍历
	grid[row][col] = 1
	if row == m-1 && col == n-1 {
		return path
	}
	successors := backtrack(grid, row+1, col)
	if successors != nil {
		return append(path, successors...)
	}
	successors = backtrack(grid, row, col+1)
	if successors != nil {
		return append(path, successors...)
	}
	return nil
}

// 朴素的回溯2
func pathWithObstacles2(obstacleGrid [][]int) [][]int {
	var ans [][]int
	dfs(obstacleGrid, &ans, make([][]int, 0, len(obstacleGrid)+len(obstacleGrid[0])-1), 0, 0)
	return ans
}

func dfs(grid [][]int, ans *[][]int, path [][]int, row, col int) {
	if *ans != nil {
		return
	}
	m, n := len(grid), len(grid[0])
	if row == m || col == n || grid[row][col] == 1 {
		return
	}
	// 防止重复遍历
	grid[row][col] = 1
	path = append(path, []int{row, col})
	if row == m-1 && col == n-1 {
		*ans = path
		return
	}
	dfs(grid, ans, path, row+1, col)
	dfs(grid, ans, path, row, col+1)
}

// 先求出终点的连通域，然后从起点出发找路径
func pathWithObstacles3(obstacleGrid [][]int) [][]int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return nil
	}
	obstacleGrid[m-1][n-1] = -1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] != -1 {
				continue
			}
			if i-1 >= 0 && obstacleGrid[i-1][j] == 0 {
				obstacleGrid[i-1][j] = -1
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				obstacleGrid[i][j-1] = -1
			}
		}
	}
	if obstacleGrid[0][0] != -1 {
		return nil
	}
	var path [][]int
	row, col := 0, 0
	for {
		path = append(path, []int{row, col})
		if row == m-1 && col == n-1 {
			break
		}
		if row+1 < m && obstacleGrid[row+1][col] == -1 {
			row++
		} else {
			col++
		}
	}
	return path
}
