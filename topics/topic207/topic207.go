package topic207

import "leetcode-go/graph"

/*
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。


提示：

1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同
通过次数126,697提交次数232,531

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/2/21 2:42 PM
*/

//dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var nodes []*graph.Node
	for _, pre := range prerequisites {
		addNode(&nodes, pre[0])
		addNode(&nodes, pre[1])
		n1 := nodes[pre[0]]
		n2 := nodes[pre[1]]
		n1.Adjacent = append(n1.Adjacent, n2)
	}
	status := make([]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		if !dfsVisit(nodes, status, i) {
			return false
		}
	}
	return true
}

func dfsVisit(nodes []*graph.Node, status []int, i int) bool {
	if status[i] == 1 {
		return false
	}
	if status[i] == 2 {
		return true
	}
	status[i] = 1
	node := nodes[i]
	for _, adj := range node.Adjacent {
		if !dfsVisit(nodes, status, adj.Val) {
			return false
		}
	}
	status[i] = 2
	return true
}

func addNode(nodes *[]*graph.Node, newNode int) {
	for i := len(*nodes); i <= newNode; i++ {
		*nodes = append(*nodes, &graph.Node{Val: i})
	}
}

// bfs
func canFinish2(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		inDegrees[p[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	finished := 0
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		for _, adjacent := range edges[node] {
			inDegrees[adjacent]--
			if inDegrees[adjacent] == 0 {
				queue = append(queue, adjacent)
			}
		}
		finished++
	}
	return finished == numCourses
}

// 原理同方法2，但是这里是通过访问的边数来判断是否连通
func canFinish3(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		inDegrees[p[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	visitedEdges := 0
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		for _, adjacent := range edges[node] {
			inDegrees[adjacent]--
			visitedEdges++
			if inDegrees[adjacent] == 0 {
				queue = append(queue, adjacent)
			}
		}
	}
	return visitedEdges == len(prerequisites)
}
