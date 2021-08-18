package topic9527

/*
自己创建的题目：判断无向图是否有环

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/3/21 8:45 PM
*/

func hasLoop(n int, edges [][]int) bool {
	adjList := make([][]int, n)
	inDeg := make([]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		adjList[i] = append(adjList[i], j)
		adjList[j] = append(adjList[j], i)
		inDeg[i]++
		inDeg[j]++
	}
	var q []int
	for i := 0; i < n; i++ {
		if inDeg[i] <= 1 {
			q = append(q, i)
		}
	}
	sum := 0
	for i := 0; i < len(q); i++ {
		node := q[i]
		for _, adj := range adjList[node] {
			if inDeg[adj] == 0 {
				continue
			}
			inDeg[adj]--
			inDeg[node]--
			sum += 2
			if inDeg[adj] == 1 {
				q = append(q, adj)
			}
		}
	}
	return sum != 2*(len(edges))
}

// 原理同上，不同点在于：统计的是访问的节点
func hasLoop2(n int, edges [][]int) bool {
	adjList := make([][]int, n)
	inDeg := make([]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		adjList[i] = append(adjList[i], j)
		adjList[j] = append(adjList[j], i)
		inDeg[i]++
		inDeg[j]++
	}
	var q []int
	for i := 0; i < n; i++ {
		if inDeg[i] <= 1 {
			q = append(q, i)
		}
	}
	visited := 0
	for i := 0; i < len(q); i++ {
		node := q[i]
		for _, adj := range adjList[node] {
			if inDeg[adj] == 0 {
				continue
			}
			inDeg[adj]--
			inDeg[node]--
			if inDeg[adj] == 1 {
				q = append(q, adj)
			}
		}
		visited++
	}
	return visited != n
}
