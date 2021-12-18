// 剑指 Offer II 105. 岛屿的最大面积
//给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
//
//一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
//找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
//
//
//
//示例 1:
//
//
//
//输入: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//输出: 6
//解释: 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
//示例 2:
//
//输入: grid = [[0,0,0,0,0,0,0,0]]
//输出: 0
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 50
//grid[i][j] is either 0 or 1
//
//
//注意：本题与主站 695 题相同： https://leetcode-cn.com/problems/max-area-of-island/
//
//通过次数5,295提交次数7,629
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/17/21 12:44 PM
package ZL6zAn

// 深度优先遍历
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var getArea func(int, int) int
	getArea = func(i, j int) int {
		if i < 0 || i == m || j < 0 || j == n {
			return 0
		}
		if grid[i][j] == 0 || visited[i][j] {
			return 0
		}
		visited[i][j] = true
		return getArea(i-1, j) + getArea(i+1, j) + getArea(i, j-1) + getArea(i, j+1) + 1
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area := getArea(i, j)
			if area > ans {
				ans = area
			}
		}
	}
	return ans
}

// 广度优先遍历
func maxAreaOfIsland2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	ans := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 || visited[i][j] {
				continue
			}
			area := 0
			q := [][2]int{{i, j}}
			visited[i][j] = true
			for len(q) > 0 {
				node := q[0]
				q = q[1:]
				area++
				for _, dir := range dirs {
					x, y := node[0]+dir[0], node[1]+dir[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && !visited[x][y] {
						q = append(q, [2]int{x, y})
						visited[x][y] = true
					}
				}
			}
			if area > ans {
				ans = area
			}
		}
	}
	return ans
}

// 带权的并查集
func maxAreaOfIsland3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	length := m * n
	parents := make([]int, length)
	ranks := make([]int, length)
	var find func(int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}
	union := func(i, j int) {
		x, y := find(i), find(j)
		if x == y {
			return
		}
		if x < y {
			x, y = y, x
		}
		//总是将索引最大的节点设置为祖先，所以后面只需要一次遍历
		parents[y] = x
		ranks[x] += ranks[y]
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			index := i*n + j
			parents[index] = index
			ranks[index] = 1
			if i-1 >= 0 && grid[i-1][j] == 1 {
				union(index, (i-1)*n+j)
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				union(index, i*n+j-1)
			}
			if ranks[index] > ans {
				ans = ranks[index]
			}
		}
	}
	return ans
}
