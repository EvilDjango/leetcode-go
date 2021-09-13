package topic0210

/*
210. 课程表 II
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。

可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

示例 1:

输入: 2, [[1,0]]
输出: [0,1]
解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
示例 2:

输入: 4, [[1,0],[2,0],[3,1],[3,2]]
输出: [0,1,2,3] or [0,2,1,3]
解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
说明:

输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
提示:

这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。

通过次数80,133提交次数148,726

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/4/21 2:35 PM
*/

// dfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	stack := make([]int, 0, numCourses)
	statuses := make([]int, numCourses)
	graph := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]bool, numCourses)
	}
	for _, pre := range prerequisites {
		graph[pre[1]][pre[0]] = true
	}
	for i := 0; i < numCourses; i++ {
		if !dfsFindOrder(graph, &stack, statuses, i) {
			return nil
		}
	}
	for i := 0; i < numCourses/2; i++ {
		stack[i], stack[numCourses-1-i] = stack[numCourses-1-i], stack[i]
	}
	return stack
}

func dfsFindOrder(graph [][]bool, stack *[]int, statuses []int, i int) bool {
	if statuses[i] == 2 {
		return true
	}
	if statuses[i] == 1 {
		return false
	}
	statuses[i] = 1
	for v, val := range graph[i] {
		if !val {
			continue
		}
		if !dfsFindOrder(graph, stack, statuses, v) {
			return false
		}
	}
	statuses[i] = 2
	*stack = append(*stack, i)
	return true
}

// 每次移除入度为1的点
func findOrder2(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for _, pre := range prerequisites {
		u, v := pre[1], pre[0]
		graph[u] = append(graph[u], v)
		inDegrees[v]++
	}
	var q []int
	for i, degree := range inDegrees {
		if degree == 0 {
			q = append(q, i)
		}
	}
	for i := 0; i < len(q); i++ {
		cur := q[i]
		for _, adj := range graph[cur] {
			inDegrees[adj]--
			if inDegrees[adj] == 0 {
				q = append(q, adj)
			}
		}
	}
	if len(q) < numCourses {
		return nil
	}
	return q
}
