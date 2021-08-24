package topic261

/*
261. 以图判树
给定从 0 到 n-1 标号的 n 个结点，和一个无向边列表（每条边以结点对来表示），请编写一个函数用来判断这些边是否能够形成一个合法有效的树结构。

示例 1：

输入: n = 5, 边列表 edges = [[0,1], [0,2], [0,3], [1,4]]
输出: true
示例 2:

输入: n = 5, 边列表 edges = [[0,1], [1,2], [2,3], [1,3], [1,4]]
输出: false
注意：你可以假定边列表 edges 中不会出现重复的边。由于所有的边是无向边，边 [0,1] 和边 [1,0] 是相同的，因此不会同时出现在边列表 edges 中。

通过次数7,495提交次数15,052

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 2:45 PM
*/

// bfs
func validTree(n int, edges [][]int) bool {
	eMap := make(map[int][]int)
	for _, edge := range edges {
		eMap[edge[0]] = append(eMap[edge[0]], edge[1])
		eMap[edge[1]] = append(eMap[edge[1]], edge[0])
	}
	count := make(map[int]bool)
	nodes := []int{0}
	prevs := []int{-1}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		prev := prevs[0]
		prevs = prevs[1:]
		if count[node] {
			return false
		}
		count[node] = true
		for _, adjacent := range eMap[node] {
			if adjacent != prev {
				nodes = append(nodes, adjacent)
				prevs = append(prevs, node)
			}
		}
	}
	return len(count) == n
}

// 并查集
func validTree2(n int, edges [][]int) bool {
	ancestors := make([]int, n)
	for i := 0; i < n; i++ {
		ancestors[i] = i
	}
	for _, edge := range edges {
		if !merge(ancestors, edge[0], edge[1]) {
			return false
		}
	}
	root := 0
	for i := 0; i < n; i++ {
		if i == ancestors[i] {
			root++
			if root > 1 {
				return false
			}
		}
	}
	return true
}

func merge(ancestors []int, i, j int) bool {
	a1 := getAncestor(ancestors, i)
	a2 := getAncestor(ancestors, j)
	if a1 == a2 {
		return false
	}
	ancestors[a2] = a1
	return true
}

func getAncestor(ancestors []int, i int) int {
	if ancestors[i] != i {
		ancestors[i] = getAncestor(ancestors, ancestors[i])
	}
	return ancestors[i]
}

// dfs
func validTree3(n int, edges [][]int) bool {
	adjacents := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjacents[u] = append(adjacents[u], v)
		adjacents[v] = append(adjacents[v], u)
	}
	statuses := make([]int, n)
	count := 0
	dfsValidTree(adjacents, statuses, -1, 0, &count)
	return count == n
}

func dfsValidTree(adjacents [][]int, statuses []int, prev, i int, count *int) bool {
	if statuses[i] == 1 {
		return false
	}
	statuses[i] = 1
	for _, adjacent := range adjacents[i] {
		if adjacent == prev {
			continue
		}
		if !dfsValidTree(adjacents, statuses, i, adjacent, count) {
			return false
		}
	}
	*count++
	return true
}
