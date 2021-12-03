// 面试题 04.01. 节点间通路
//节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
//
//示例1:
//
// 输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
// 输出：true
//示例2:
//
// 输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
// 输出 true
//提示：
//
//节点数量n在[0, 1e5]范围内。
//节点编号大于等于 0 小于 n。
//图中可能存在自环和平行边。
//通过次数18,811提交次数35,439
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/3/21 2:39 PM
package route_between_nodes_lcci

// 邻接表，广度遍历
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	adjacentList := make([][]int, n)
	for _, link := range graph {
		src, dst := link[0], link[1]
		adjacentList[src] = append(adjacentList[src], dst)
	}
	q := []int{start}
	visited := make([]bool, n)
	for len(q) > 0 {
		node := q[0]
		visited[node] = true
		q = q[1:]
		for _, adj := range adjacentList[node] {
			if adj == target {
				return true
			}
			if !visited[adj] {
				q = append(q, adj)
			}
		}
	}
	return false
}

// 邻接表，深度遍历
func findWhetherExistsPath2(n int, graph [][]int, start int, target int) bool {
	adjacentList := make([][]int, n)
	for _, link := range graph {
		src, dst := link[0], link[1]
		adjacentList[src] = append(adjacentList[src], dst)
	}
	visited := make([]bool, n)
	var dfs func(src int) bool
	dfs = func(src int) bool {
		visited[src] = true
		for _, adj := range adjacentList[src] {
			if adj == target {
				return true
			}
			if !visited[adj] && dfs(adj) {
				return true
			}
		}
		return false
	}
	return dfs(start)
}
