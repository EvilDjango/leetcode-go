// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/20 上午11:18
package minimum_moves_to_spread_stones_over_grid

import (
	"fmt"
)

func minimumMoves(grid [][]int) int {
	var riches, poors [][]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] > 1 {
				riches = append(riches, []int{i, j})
			} else if grid[i][j] == 0 {
				poors = append(poors, []int{i, j})
			}
		}
	}
	if len(poors) == 0 {
		return 0
	}

	getStatus := func() string {
		return fmt.Sprint(grid)
	}
	cache := make(map[string]int)

	var dfs func(int) int
	dfs = func(poorNodes int) int {
		if poorNodes == 0 {
			return 0
		}
		status := getStatus()
		if moves, ok := cache[status]; ok {
			return moves
		}

		moves := 1000
		for _, rich := range riches {
			row, col := rich[0], rich[1]
			if grid[row][col] == 1 {
				continue
			}
			grid[row][col]--
			for _, poor := range poors {
				pRow, pCol := poor[0], poor[1]
				if grid[pRow][pCol] != 0 {
					continue
				}
				dis := distance(row, col, pRow, pCol)
				grid[pRow][pCol] = 1
				moves = min(moves, dfs(poorNodes-1)+dis)
				grid[pRow][pCol] = 0
			}
			grid[row][col]++
		}
		cache[status] = moves
		return moves
	}
	return dfs(len(poors))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func distance(row1, col1, row2, col2 int) int {
	return abs(row1-row2) + abs(col1-col2)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
