// 假设你设计一个游戏，用一个 m 行 n 列的 2D 网格来存储你的游戏地图。
//
//起始的时候，每个格子的地形都被默认标记为「水」。我们可以通过使用 addLand 进行操作，将位置 (row, col) 的「水」变成「陆地」。
//
//你将会被给定一个列表，来记录所有需要被操作的位置，然后你需要返回计算出来 每次 addLand 操作后岛屿的数量。
//
//注意：一个岛的定义是被「水」包围的「陆地」，通过水平方向或者垂直方向上相邻的陆地连接而成。你可以假设地图网格的四边均被无边无际的「水」所包围。
//
//请仔细阅读下方示例与解析，更加深入了解岛屿的判定。
//
//示例:
//
//输入: m = 3, n = 3, positions = [[0,0], [0,1], [1,2], [2,1]]
//输出: [1,1,2,3]
//解析:
//
//起初，二维网格 grid 被全部注入「水」。（0 代表「水」，1 代表「陆地」）
//
//0 0 0
//0 0 0
//0 0 0
//操作 #1：addLand(0, 0) 将 grid[0][0] 的水变为陆地。
//
//1 0 0
//0 0 0   Number of islands = 1
//0 0 0
//操作 #2：addLand(0, 1) 将 grid[0][1] 的水变为陆地。
//
//1 1 0
//0 0 0   岛屿的数量为 1
//0 0 0
//操作 #3：addLand(1, 2) 将 grid[1][2] 的水变为陆地。
//
//1 1 0
//0 0 1   岛屿的数量为 2
//0 0 0
//操作 #4：addLand(2, 1) 将 grid[2][1] 的水变为陆地。
//
//1 1 0
//0 0 1   岛屿的数量为 3
//0 1 0
//拓展：
//
//你是否能在 O(k log mn) 的时间复杂度程度内完成每次的计算？（k 表示 positions 的长度）
//
//通过次数4,010提交次数10,646
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-islands-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// @author xuejunc deerhunter0837@gmail.com
// @create 8/31/21 10:18 AM
package topic305

import (
	"fmt"
)

func numIslands2(m int, n int, positions [][]int) []int {
	size := len(positions)
	if size == 0 {
		return nil
	}
	unions := make(map[string]string, size)
	ans := make([]int, size)
	for i, position := range positions {
		if i > 0 {
			ans[i] = ans[i-1]
		}
		key1 := fmt.Sprintf("%d-%d", position[0], position[1])
		//忽略重复操作
		if unions[key1] != "" {
			continue
		}
		ans[i]++
		unions[key1] = key1
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, direct := range directions {
			x, y := position[0]+direct[0], position[1]+direct[1]
			if x < 0 || y < 0 || x == m || y == n {
				continue
			}
			key2 := fmt.Sprintf("%d-%d", x, y)
			if unions[key2] != "" && merge(unions, key1, key2) {
				ans[i]--
			}
		}
	}
	return ans
}

func merge(unions map[string]string, key1, key2 string) bool {
	root1 := getRoot(unions, key1)
	root2 := getRoot(unions, key2)
	if root1 == root2 {
		return false
	}
	unions[root1] = root2
	return true
}

func getRoot(unions map[string]string, key string) string {
	// 当一个键没有祖先时，它自己就是自己的祖先
	if unions[key] != key {
		unions[key] = getRoot(unions, unions[key])
	}
	return unions[key]
}
