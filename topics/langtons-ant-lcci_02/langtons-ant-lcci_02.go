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
package langtons_ant_lcci_02

import (
	"strings"
)

func printKMoves(K int) []string {
	x, y := 0, 0
	left, right, top, down := 0, 0, 0, 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	names := []byte{'R', 'D', 'L', 'U'}
	direct := 0
	nodes := map[int]map[int]byte{}
	nodes[0] = map[int]byte{0: '_'}
	for i := 0; i <= K; i++ {
		if i == K {
			nodes[x][y] = names[direct]
			continue
		}
		if nodes[x][y] == '_' {
			nodes[x][y] = 'X'
			direct++
		} else {
			nodes[x][y] = '_'
			direct--
		}
		direct &= 3
		x += directions[direct][0]
		y += directions[direct][1]
		left = Min(left, y)
		right = Max(right, y)
		top = Min(top, x)
		down = Max(down, x)
		if _, ok := nodes[x]; !ok {
			nodes[x] = map[int]byte{}
		}
		if _, ok := nodes[x][y]; !ok {
			nodes[x][y] = '_'
		}
	}
	ans := make([]string, down-top+1)
	for i := top; i <= down; i++ {
		builder := strings.Builder{}
		for j := left; j <= right; j++ {
			if b, ok := nodes[i][j]; ok {
				builder.WriteByte(b)
			} else {
				builder.WriteByte('_')
			}
		}
		ans[i-top] = builder.String()
	}
	return ans
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//参考优秀题解做了优化，优化点有两个：
//第一，使用一个long的前半段记录row，后半段记录col，从而使用一层map即可
//第二，仅记录黑色网格，节省空间
func printKMoves2(K int) []string {
	// 我们不能让x,y为负值，因为负值可能会占据int64超过一半的位，故x,y都初始化为K的最大值
	x, y := int64(100000), int64(100000)
	left, right, top, down := y, y, x, x
	// 这里是4个方向，往后是顺时针变换，往前是逆时针变换
	directions := [][]int64{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direct := 0
	blacks := map[int64]bool{}
	for i := 0; i < K; i++ {
		key := x<<20 | y
		if blacks[key] {
			delete(blacks, key)
			direct--
		} else {
			blacks[key] = true
			direct++
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
		builder := strings.Builder{}
		for j := left; j <= right; j++ {
			// 最后一个位置需要指明蚂蚁最后的方向
			if i == x && j == y {
				builder.WriteByte([]byte{'R', 'D', 'L', 'U'}[direct])
				continue
			}
			if _, ok := blacks[i<<20|j]; ok {
				builder.WriteByte('X')
			} else {
				builder.WriteByte('_')
			}
		}
		ans[i-top] = builder.String()
	}
	return ans
}

func min(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
