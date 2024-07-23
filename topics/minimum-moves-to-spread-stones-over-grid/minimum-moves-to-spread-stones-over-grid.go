// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/20 上午11:18
package minimum_moves_to_spread_stones_over_grid

import (
	"fmt"
	"math"
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

func minimumMoves2(grid [][]int) int {
	var riches, poors [][]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 1; k < grid[i][j]; k++ {
				riches = append(riches, []int{i, j})
			}
			if grid[i][j] == 0 {
				poors = append(poors, []int{i, j})
			}
		}
	}
	if len(poors) == 0 {
		return 0
	}

	ans := math.MaxInt32
	for {
		steps := 0
		for i, rich := range riches {
			steps += dis(rich, poors[i])
		}
		if steps < ans {
			ans = steps
		}
		if !nextPermutation(riches) {
			break
		}
	}
	return ans
}

func nextPermutation(riches [][]int) bool {
	n := len(riches)
	p := -1
	for i := 0; i < n-1; i++ {
		if less(riches[i], riches[i+1]) {
			p = i
		}
	}
	if p == -1 {
		return false
	}
	q := p + 1
	for i := q + 1; i < n; i++ {
		if less(riches[p], riches[i]) {
			q = i
		}
	}
	riches[p], riches[q] = riches[q], riches[p]
	left, right := p+1, n-1
	for left < right {
		riches[left], riches[right] = riches[right], riches[left]
		left++
		right--
	}
	return true
}

func less(a, b []int) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

func dis(i, j []int) int {
	return abs(i[0]-j[0]) + abs(i[1]-j[1])
}
