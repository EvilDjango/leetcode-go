// 面试题 17.23. 最大黑方阵
//给定一个方阵，其中每个单元(像素)非黑即白。设计一个算法，找出 4 条边皆为黑色像素的最大子方阵。
//
//返回一个数组 [r, c, size] ，其中 r, c 分别代表子方阵左上角的行号和列号，size 是子方阵的边长。若有多个满足条件的子方阵，返回 r 最小的，若 r 相同，返回 c 最小的子方阵。若无满足条件的子方阵，返回空数组。
//
//示例 1:
//
//输入:
//[
//   [1,0,1],
//   [0,0,1],
//   [0,0,1]
//]
//输出: [1,0,2]
//解释: 输入中 0 代表黑色，1 代表白色，标粗的元素即为满足条件的最大子方阵
//示例 2:
//
//输入:
//[
//   [0,1,1],
//   [1,0,1],
//   [1,1,0]
//]
//输出: [0,0,1]
//提示：
//
//matrix.length == matrix[0].length <= 200
//通过次数4,351提交次数11,471
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/13/21 3:51 PM
package topic2339

// 分别计算每个位置往右往下的最长连续黑色序列
func findSquare(matrix [][]int) []int {
	var ans []int
	m := len(matrix)
	if m == 0 {
		return ans
	}
	n := len(matrix[0])

	// rightSeries[i][j][表示从位置(i,j)向右的最长连续黑色序列长度，
	rightSeries := make([][]int, m)
	// downSeries[i][j]表示从位置(i,j)向下的最长连续黑色序列长度
	downSeries := make([][]int, m)
	// downSeries[i][j]表示从位置(i,j)向左的最长连续黑色序列长度
	leftSeries := make([][]int, m)
	// downSeries[i][j]表示从位置(i,j)向上的最长连续黑色序列长度
	upSeries := make([][]int, m)
	for i := 0; i < m; i++ {
		rightSeries[i] = make([]int, n)
		downSeries[i] = make([]int, n)
		leftSeries[i] = make([]int, n)
		upSeries[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		sum := 0
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 1 {
				sum = 0
			} else {
				sum++
				rightSeries[i][j] = sum
			}
		}
	}

	for i := 0; i < n; i++ {
		sum := 0
		for j := m - 1; j >= 0; j-- {
			if matrix[j][i] == 1 {
				sum = 0
			} else {
				sum++
				downSeries[j][i] = sum
			}
		}
	}

	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				sum = 0
			} else {
				sum++
				leftSeries[i][j] = sum
			}
		}
	}

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			if matrix[j][i] == 1 {
				sum = 0
			} else {
				sum++
				upSeries[j][i] = sum
			}
		}
	}

	maxSize := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			possibleSize := min(rightSeries[i][j], downSeries[i][j])
			size := biggerSquare(i, j, possibleSize, maxSize, leftSeries, upSeries)
			if size != -1 {
				ans = []int{i, j, size}
				maxSize = size
			}
		}
	}
	return ans
}

func biggerSquare(row, col, possibleSize, maxSize int, leftSeries, upSeries [][]int) int {
	for ; possibleSize > maxSize; possibleSize-- {
		r := row + possibleSize - 1
		c := col + possibleSize - 1
		if leftSeries[r][c] >= possibleSize && upSeries[r][c] >= possibleSize {
			return possibleSize
		}
	}
	return -1
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
