package topic286

/*
286. 墙与门
你被给定一个 m × n 的二维网格 rooms ，网格中有以下三种可能的初始化值：

-1 表示墙或是障碍物
0 表示一扇门
INF 无限表示一个空的房间。然后，我们用 231 - 1 = 2147483647 代表 INF。你可以认为通往门的距离总是小于 2147483647 的。
你要给每个空房间位上填上该房间到 最近门的距离 ，如果无法到达门，则填 INF 即可。



示例 1：


输入：rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
输出：[[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]
示例 2：

输入：rooms = [[-1]]
输出：[[-1]]
示例 3：

输入：rooms = [[2147483647]]
输出：[[2147483647]]
示例 4：

输入：rooms = [[0]]
输出：[[0]]


提示：

m == rooms.length
n == rooms[i].length
1 <= m, n <= 250
rooms[i][j] 是 -1、0 或 231 - 1
通过次数10,795提交次数21,344

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/13/21 1:12 PM
*/

// dfs
func wallsAndGates(rooms [][]int) {
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				dfs(rooms, i, j+1, 1)
				dfs(rooms, i, j-1, 1)
				dfs(rooms, i+1, j, 1)
				dfs(rooms, i-1, j, 1)
			}
		}

	}
}

func dfs(rooms [][]int, i, j, distance int) {
	if i == -1 || j == -1 || i == len(rooms) || j == len(rooms[0]) {
		return
	}
	val := rooms[i][j]
	if val == 0 || val == -1 || val <= distance {
		return
	}
	rooms[i][j] = distance
	distance++
	dfs(rooms, i, j+1, distance)
	dfs(rooms, i, j-1, distance)
	dfs(rooms, i+1, j, distance)
	dfs(rooms, i-1, j, distance)
}

//bfs
func wallsAndGates2(rooms [][]int) {
	var q [][]int
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(q); i++ {
		p := q[i]
		for _, direct := range directions {
			j, k := p[0]+direct[0], p[1]+direct[1]
			if j < 0 || j == len(rooms) || k < 0 || k == len(rooms[0]) {
				continue
			}

			val := rooms[j][k]
			if val == -1 || val == 0 || val <= rooms[p[0]][p[1]]+1 {
				continue
			}
			rooms[j][k] = rooms[p[0]][p[1]] + 1
			q = append(q, []int{j, k})
		}
	}
}
