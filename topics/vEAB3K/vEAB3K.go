// 剑指 Offer II 106. 二分图
//存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。
//
//给定一个二维数组 graph ，表示图，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下属性：
//
//不存在自环（graph[u] 不包含 u）。
//不存在平行边（graph[u] 不包含重复值）。
//如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
//这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
//二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。
//
//如果图是二分图，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//
//
//输入：graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
//输出：false
//解释：不能将节点分割成两个独立的子集，以使每条边都连通一个子集中的一个节点与另一个子集中的一个节点。
//示例 2：
//
//
//
//输入：graph = [[1,3],[0,2],[1,3],[0,2]]
//输出：true
//解释：可以将节点分成两组: {0, 2} 和 {1, 3} 。
//
//
//提示：
//
//graph.length == n
//1 <= n <= 100
//0 <= graph[u].length < n
//0 <= graph[u][i] <= n - 1
//graph[u] 不会包含 u
//graph[u] 的所有值 互不相同
//如果 graph[u] 包含 v，那么 graph[v] 也会包含 u
//
//
//注意：本题与主站 785 题相同： https://leetcode-cn.com/problems/is-graph-bipartite/
//
//通过次数2,029提交次数3,669
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/16/21 1:51 PM
package vEAB3K

// 笨办法，老打老实地回溯
// 一个关键点：同一个集合中的点之间不能存在边,或者说，相邻的点不能在同一个集合中
func isBipartite(graph [][]int) bool {
	n := len(graph)
	groups := make([]int, n)
	for i := 0; i < n; i++ {
		groups[i] = -1
	}
	groupAdjacent := func(i, group int, groups []int) {
		for _, node := range graph[i] {
			groups[node] = group
		}
	}
	var dfs func(int, []int) bool
	dfs = func(i int, groups []int) bool {
		if i == n {
			return true
		}
		// 孤立的点放在哪边都无所谓
		if len(graph[i]) == 0 {
			return dfs(i+1, groups)
		}
		temp := make([]int, n)
		copy(temp, groups)
		groups = temp

		adjacentGroup := -1
		for _, node := range graph[i] {
			if groups[node] != -1 {
				// 同一个点的相邻点应该全部在同一个集合
				if adjacentGroup != -1 && adjacentGroup != groups[node] {
					return false
				}
				adjacentGroup = groups[node]
			}
		}
		if adjacentGroup != -1 {
			groupAdjacent(i, adjacentGroup, groups)
		}
		group := groups[i]
		if adjacentGroup == group && adjacentGroup != -1 {
			return false
		}

		if adjacentGroup != -1 && group != -1 && adjacentGroup != group {
			return dfs(i+1, groups)
		}
		if adjacentGroup == -1 && group != -1 {
			groupAdjacent(i, 1-group, groups)
			return dfs(i+1, groups)
		}
		if adjacentGroup != -1 && group == -1 {
			groups[i] = 1 - adjacentGroup
			return dfs(i+1, groups)
		}

		// 下面是两者都为-1的情形,节点i和其相邻节点的分组可以为 0，1或者1，0
		groupAdjacent(i, 0, groups)
		groups[i] = 1
		if dfs(i+1, groups) {
			return true
		}
		groupAdjacent(i, 1, groups)
		groups[i] = 0
		return dfs(i+1, groups)
	}
	return dfs(0, groups)
}

// 深度遍历
func isBipartite2(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	var dfs func(int, int) bool
	dfs = func(i, color int) bool {
		if colors[i] != 0 {
			return colors[i] == color
		}
		colors[i] = color
		nextColor := color%2 + 1
		for _, neighbor := range graph[i] {
			if !dfs(neighbor, nextColor) {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

// 广度遍历
func isBipartite3(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		if colors[i] != 0 {
			continue
		}
		q := []int{i}
		colors[i] = 1
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			nextColor := colors[node]%2 + 1
			for _, neighbor := range graph[node] {
				if colors[neighbor] != 0 {
					if colors[neighbor] != nextColor {
						return false
					}
					continue
				}
				colors[neighbor] = nextColor
				q = append(q, neighbor)
			}
		}
	}
	return true
}

// 并查集
func isBipartite4(graph [][]int) bool {
	n := len(graph)
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		ranks[i] = 1
	}
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
		if ranks[x] > ranks[y] {
			x, y = y, x
		}
		parents[x] = y
		ranks[y] += ranks[x]
	}
	for i := 0; i < n; i++ {
		for _, neighbor := range graph[i] {
			if find(i) == find(neighbor) {
				return false
			}
			union(neighbor, graph[i][0])
		}
	}
	return true
}
