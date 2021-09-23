// 面试题 16.19. 水域大小
//你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
//
//示例：
//
//输入：
//[
//  [0,2,1,0],
//  [0,1,0,1],
//  [1,1,0,1],
//  [0,1,0,1]
//]
//输出： [1,2,4]
//提示：
//
//0 < len(land) <= 1000
//0 < len(land[i]) <= 1000
//通过次数20,141提交次数32,966
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/23/21 2:18 PM
package pond_sizes_lcci

// 朴素的广度遍历
func pondSizes(land [][]int) []int {
	m := len(land)
	if m == 0 {
		return nil
	}
	n := len(land[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	bfs := func(i int, j int) int {
		if visited[i][j] || land[i][j] != 0 {
			return 0
		}
		q := [][]int{{i, j}}
		visited[i][j] = true
		for i := 0; i < len(q); i++ {
			pos := q[i]
			x, y := pos[0], pos[1]
			for _, direction := range directions {
				k, l := x+direction[0], y+direction[1]
				if k >= 0 && l >= 0 && k < m && l < n && !visited[k][l] && land[k][l] == 0 {
					q = append(q, []int{k, l})
					visited[k][l] = true
				}
			}
		}
		return len(q)
	}
	sizes := make([]int, m*n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sizes[bfs(i, j)]++
		}
	}
	var ans []int
	for i := 1; i <= m*n; i++ {
		for sizes[i] > 0 {
			ans = append(ans, i)
			sizes[i]--
		}
	}
	return ans
}

// 使用并查集和基数排序
func pondSizes2(land [][]int) []int {
	m := len(land)
	if m == 0 {
		return nil
	}
	n := len(land[0])
	parents := make([]int, m*n)
	weights := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			parents[i*n+j] = i*n + j
			if land[i][j] == 0 {
				weights[i*n+j] = 1
			}
		}
	}
	// 每个点又8个相邻节点，每个节点只负责其中4个，因为对称性，剩余的相邻节点会主动来处理相邻关系
	adjacents := [][]int{{-1, 1}, {0, 1}, {1, 0}, {1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] != 0 {
				continue
			}
			for _, adj := range adjacents {
				x, y := i+adj[0], j+adj[1]
				if x >= 0 && y >= 0 && x < m && y < n && land[x][y] == 0 {
					merge(parents, weights, i*n+j, x*n+y)
				}
			}
		}
	}
	sizes := make([]int, m*n+1)
	for i := 0; i < m*n; i++ {
		if parents[i] == i {
			sizes[weights[i]]++
		}
	}
	var ans []int
	for i := 1; i <= m*n; i++ {
		for sizes[i] > 0 {
			ans = append(ans, i)
			sizes[i]--
		}
	}
	return ans
}

func merge(parents []int, weights []int, i, j int) {
	root1 := find(parents, i)
	root2 := find(parents, j)
	if root1 == root2 {
		return
	}
	if root1 < root2 {
		parents[root2] = root1
		weights[root1] += weights[root2]
	} else {
		parents[root1] = root2
		weights[root2] += weights[root1]
	}
}

func find(parents []int, i int) int {
	for parents[i] != i {
		parents[i] = parents[parents[i]]
		i = parents[i]
	}
	return i
}
