// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/3 上午12:02
package divide_nodes_into_the_maximum_number_of_groups

func magnificentSets(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	layer := 0
	layers := make([]int, n)
	groups := make([]int, n)
	groupCnt := 0
	bfs := func(start int) int {
		var group int
		if groups[start] > 0 {
			group = groups[start]
		} else {
			groupCnt++
			group = groupCnt
			groups[start] = group
		}
		startLayer := layer + 1
		q := []int{start}
		layers[start] = layer + 1
		for len(q) > 0 {
			layer++
			size := len(q)
			for i := 0; i < size; i++ {
				u := q[i]
				for _, v := range g[u] {
					if layers[v] < startLayer {
						layers[v] = layer + 1
						q = append(q, v)
						groups[v] = group
					} else if layers[v] == layer {
						return -1
					}
				}
			}
			q = q[size:]
		}
		return layer - startLayer + 1
	}
	groupMax := make(map[int]int, n)
	for i := 0; i < n; i++ {
		maxLength := bfs(i)
		if maxLength == -1 {
			return -1
		}
		group := groups[i]
		groupMax[group] = max(groupMax[group], maxLength)
	}
	ans := 0
	for _, v := range groupMax {
		ans += v
	}
	return ans
}
