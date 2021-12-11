// 剑指 Offer II 112. 最长递增路径
//给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
//
//
//
//示例 1：
//
//
//
//输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
//输出：4
//解释：最长递增路径为 [1, 2, 6, 9]。
//示例 2：
//
//
//
//输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
//输出：4
//解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
//示例 3：
//
//输入：matrix = [[1]]
//输出：1
//
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 200
//0 <= matrix[i][j] <= 231 - 1
//
//
//注意：本题与主站 329 题相同： https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/
//
//通过次数1,944提交次数3,411
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/11/21 7:13 PM
package fpTFWP

// 记忆化回溯
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	var backtrack func(int, int) int
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	backtrack = func(x, y int) int {
		if cache[x][y] != 0 {
			return cache[x][y]
		}
		deepest := 0
		for _, direct := range directions {
			r, c := x+direct[0], y+direct[1]
			if r < 0 || r == m || c < 0 || c == n || matrix[r][c] >= matrix[x][y] {
				continue
			}
			depth := backtrack(r, c)
			if depth > deepest {
				deepest = depth
			}
		}
		cache[x][y] = deepest + 1
		return deepest + 1
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans < backtrack(i, j) {
				ans = backtrack(i, j)
			}
		}
	}
	return ans
}

// 拓扑排序。把每个位置P(x,y)看成一个节点，如果P的值小于于其相邻点Q，那么我们认为存在一条有向边PQ，这样整个矩阵就转换成了一个有向图。将所有入度为零的点加入队列，
// 然后把所有队列中的点相关的边移除，得到下一批入度为0的点，继续进行下一轮循环。循环的次数记为最长递增序列的长度。
func longestIncreasingPath2(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	inDegree := make([][]int, m)
	for i := 0; i < m; i++ {
		inDegree[i] = make([]int, n)
	}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if x < 0 || x == m || y < 0 || y == n || matrix[x][y] >= matrix[i][j] {
					continue
				}
				inDegree[x][y]++
			}
		}
	}
	var q [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if inDegree[i][j] == 0 {
				q = append(q, [2]int{i, j})
			}
		}
	}
	maxLen := 0
	for len(q) > 0 {
		maxLen++
		size := len(q)
		for i := 0; i < size; i++ {
			p := q[i]
			for _, dir := range dirs {
				x, y := p[0]+dir[0], p[1]+dir[1]
				if x < 0 || x == m || y < 0 || y == n || matrix[x][y] >= matrix[p[0]][p[1]] {
					continue
				}
				inDegree[x][y]--
				if inDegree[x][y] == 0 {
					q = append(q, [2]int{x, y})
				}
			}
		}
		q = q[size:]
	}
	return maxLen
}
