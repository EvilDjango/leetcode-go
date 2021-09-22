// 面试题 16.22. 兰顿蚂蚁
//一只蚂蚁坐在由白色和黑色方格构成的无限网格上。开始时，网格全白，蚂蚁面向右侧。每行走一步，蚂蚁执行以下操作。
//
//(1) 如果在白色方格上，则翻转方格的颜色，向右(顺时针)转 90 度，并向前移动一个单位。
//(2) 如果在黑色方格上，则翻转方格的颜色，向左(逆时针方向)转 90 度，并向前移动一个单位。
//
//编写程序来模拟蚂蚁执行的前 K 个动作，并返回最终的网格。
//
//网格由数组表示，每个元素是一个字符串，代表网格中的一行，黑色方格由 'X' 表示，白色方格由 '_' 表示，蚂蚁所在的位置由 'L', 'U', 'R', 'D' 表示，分别表示蚂蚁 左、上、右、下 的朝向。只需要返回能够包含蚂蚁走过的所有方格的最小矩形。
//
//示例 1:
//
//输入: 0
//输出: ["R"]
//示例 2:
//
//输入: 2
//输出:
//[
//  "_X",
//  "LX"
//]
//示例 3:
//
//输入: 5
//输出:
//[
//  "_U",
//  "X_",
//  "XX"
//]
//说明：
//
//K <= 100000
//通过次数2,677提交次数4,594
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/22/21 7:36 PM
package langtons_ant_lcci

// 直接分配一个足够大的二维数组，会超出内存限制
func printKMoves(K int) []string {
	if K == 0 {
		return []string{"R"}
	}
	size := K<<1 + 1
	matrix := make([][]byte, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = '_'
		}
	}
	x, y := K, K
	left, right, top, down := K, K, K, K
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	names := []byte{'R', 'D', 'L', 'U'}
	direct := 0
	for i := 0; i <= K; i++ {
		if i == K {
			matrix[x][y] = names[direct]
			continue
		}
		if matrix[x][y] == '_' {
			matrix[x][y] = 'X'
			direct++
		} else {
			matrix[x][y] = '_'
			direct--
		}
		direct &= 3
		x += directions[direct][0]
		y += directions[direct][1]
		left = min(left, y)
		right = max(right, y)
		top = min(top, x)
		down = max(down, x)
	}
	ans := make([]string, down-top+1)
	for i := top; i <= down; i++ {
		ans[i-top] = string(matrix[i][left : right+1])
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
