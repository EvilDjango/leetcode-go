package topic302

/*
302. 包含全部黑色像素的最小矩形
图片在计算机处理中往往是使用二维矩阵来表示的。

给你一个二进制矩阵 image 表示一张黑白图片，0 代表白色像素，1 代表黑色像素。

黑色像素相互连接，也就是说，图片中只会有一片连在一块儿的黑色像素。像素点是水平或竖直方向连接的。

给你两个整数 x 和 y 表示某一个黑色像素的位置，请你找出包含全部黑色像素的最小矩形（与坐标轴对齐），并返回该矩形的面积。





示例 1：


输入：image = [["0","0","1","0"],["0","1","1","0"],["0","1","0","0"]], x = 0, y = 2
输出：6
示例 2：

输入：image = [["1"]], x = 0, y = 0
输出：1


提示：

m == image.length
n == image[i].length
1 <= m, n <= 100
image[i][j] 为 '0' 或 '1'
1 <= x < m
1 <= y < n
image[x][y] == '1'.
image 中的黑色像素仅形成一个 组件
通过次数1,855提交次数2,760

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/24/21 5:16 PM
*/
// bfs 会超过空间限制
func minArea(image [][]byte, x int, y int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(image), len(image[0])
	xMin, xMax, yMin, yMax := m, 0, n, 0
	q := [][]int{{x, y}}
	for i := 0; i < len(q); i++ {
		node := q[i]
		x, y := node[0], node[1]
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
		image[x][y] = '3'
		for _, direct := range directions {
			x1, y1 := x+direct[0], y+direct[1]
			if x1 < 0 || y1 < 0 || x1 == m || y1 == n {
				continue
			}
			if image[x1][y1] == '1' {
				q = append(q, []int{x1, y1})
			}
		}
	}

	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if image[i][j] == '3' {
				image[i][j] = '1'
			}
		}
	}
	return (xMax - xMin + 1) * (yMax - yMin + 1)
}

// dfs
func minArea2(image [][]byte, x int, y int) int {
	m, n := len(image), len(image[0])
	xMin, xMax, yMin, yMax := m, 0, n, 0
	var dfs func(i, j int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x == m || y == n || image[x][y] != '1' {
			return
		}
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
		image[x][y] = '0'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	dfs(x, y)
	return (xMax - xMin + 1) * (yMax - yMin + 1)
}

// 二分查找
func minArea3(image [][]byte, x int, y int) int {
	m, n := len(image), len(image[0])
	xMin := findXExtremum(image, 0, x, true)
	xMax := findXExtremum(image, x, m-1, false)
	yMin := findYExtremum(image, 0, y, true)
	yMax := findYExtremum(image, y, n-1, false)
	return (xMax - xMin + 1) * (yMax - yMin + 1)
}

func findXExtremum(image [][]byte, l, r int, min bool) int {
	for l <= r {
		mid := (r-l)/2 + l
		found := false
		for _, b := range image[mid] {
			if b == '1' {
				found = true
				break
			}
		}
		if found == min {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if min {
		return l
	}
	return r
}

func findYExtremum(image [][]byte, l, r int, min bool) int {
	for l <= r {
		mid := (r-l)/2 + l
		found := false
		for _, row := range image {
			if row[mid] == '1' {
				found = true
				break
			}
		}
		if found == min {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if min {
		return l
	}
	return r
}
