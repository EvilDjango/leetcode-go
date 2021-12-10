// 剑指 Offer II 116. 省份数量
//有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
//
//省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
//
//给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
//
//返回矩阵中 省份 的数量。
//
//
//
//示例 1：
//
//
//输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
//输出：2
//示例 2：
//
//
//输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
//输出：3
//
//
//提示：
//
//1 <= n <= 200
//n == isConnected.length
//n == isConnected[i].length
//isConnected[i][j] 为 1 或 0
//isConnected[i][i] == 1
//isConnected[i][j] == isConnected[j][i]
//
//
//注意：本题与主站 547 题相同： https://leetcode-cn.com/problems/number-of-provinces/
//
//通过次数4,936提交次数7,615
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/10/21 9:49 PM
package bLyHh0

// 并查集
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(i, j int) {
		x, y := find(i), find(j)
		parent[x] = y
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}
	return count
}

// 深度优先搜索
func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	var dfs func(int) bool
	dfs = func(i int) bool {
		if visited[i] {
			return false
		}
		visited[i] = true
		for j, val := range isConnected[i] {
			if val == 1 {
				dfs(j)
			}
		}
		return true
	}
	count := 0
	for i := 0; i < n; i++ {
		if dfs(i) {
			count++
		}
	}
	return count
}

// 广度优先搜索
func findCircleNum3(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		count++
		q := []int{i}
		for len(q) > 0 {
			j := q[0]
			q = q[1:]
			for k, val := range isConnected[j] {
				if !visited[k] && val == 1 {
					q = append(q, k)
					visited[j] = true
				}
			}
		}
	}
	return count
}
