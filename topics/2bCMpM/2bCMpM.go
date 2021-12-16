// 剑指 Offer II 107. 矩阵中的距离
//给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
//
//两个相邻元素间的距离为 1 。
//
//
//
//示例 1：
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
//输出：[[0,0,0],[0,1,0],[0,0,0]]
//示例 2：
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
//输出：[[0,0,0],[0,1,0],[1,2,1]]
//
//
//提示：
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 104
//1 <= m * n <= 104
//mat[i][j] is either 0 or 1.
//mat 中至少有一个 0
//
//
//注意：本题与主站 542 题相同：https://leetcode-cn.com/problems/01-matrix/
//
//通过次数2,022提交次数3,696
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/16/21 1:26 PM
package _bCMpM

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	var q [][2]int
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				ans[i][j] = 0
				q = append(q, [2]int{i, j})
			} else {
				ans[i][j] = -1
			}
		}
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r >= 0 && r < m && c >= 0 && c < n && ans[r][c] == -1 {
				ans[r][c] = ans[x][y] + 1
				q = append(q, [2]int{r, c})
			}
		}
	}
	return ans
}
