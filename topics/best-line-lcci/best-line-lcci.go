// 面试题 16.14. 最佳直线
//给定一个二维平面及平面上的 N 个点列表Points，其中第i个点的坐标为Points[i]=[Xi,Yi]。请找出一条直线，其通过的点的数目最多。
//
//设穿过最多点的直线所穿过的全部点编号从小到大排序的列表为S，你仅需返回[S[0],S[1]]作为答案，若有多条直线穿过了相同数量的点，则选择S[0]值较小的直线返回，S[0]相同则选择S[1]值较小的直线返回。
//
//示例：
//
//输入： [[0,0],[1,1],[1,0],[2,0]]
//输出： [0,2]
//解释： 所求直线穿过的3个点的编号为[0,2,3]
//提示：
//
//2 <= len(Points) <= 300
//len(Points[i]) = 2
//通过次数3,500提交次数6,375
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/25/21 9:00 PM
package best_line_lcci

import (
	"fmt"
	"strconv"
)

// 利用直线的斜截式方程，y=kx+b，那么我们可以用字符串 k-b来表示一条直线，特殊地，对于垂直于x轴的直线，直接用其横截距来表示。为了避免浮点数精度带
//来的误判，我们将k和b表示为不可约分的m/n, 那么b可以表示为(ny-mx)/n,约分得到 p/q
func bestLine(points [][]int) []int {
	lines := map[string][]int{}
	for i := 0; i < len(points); i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			var key string
			x2, y2 := points[j][0], points[j][1]
			if x1 == x2 {
				key = strconv.Itoa(x1)
			} else if y1 == y2 {
				key = fmt.Sprintf("0-%d", y1)
			} else {
				m, n := reduce(y2-y1, x2-x1)
				p, q := reduce(n*y1-m*x1, n)
				key = fmt.Sprintf("%d/%d-%d/%d", m, n, p, q)
			}
			if lines[key] == nil {
				lines[key] = append(lines[key], i, j)
			} else if lines[key][len(lines[key])-1] < j {
				lines[key] = append(lines[key], j)
			}
		}
	}
	var ans []int
	for _, points := range lines {
		if len(points) > len(ans) || (len(points) == len(ans) && (less(points, ans))) {
			ans = points
		}
	}
	return ans[:2]
}

func less(points []int, ans []int) bool {
	if points[0] != ans[0] {
		return points[0] < ans[0]
	}
	return points[1] < ans[1]
}

func reduce(i, j int) (int, int) {
	g := gcd(i, j)
	return i / g, j / g
}
func gcd(i, j int) int {
	for i%j != 0 {
		i, j = j, i%j
	}
	return j
}

// 利用两点式表示过（x1,y1）和(x2,y2)的直线　(y-y1)*(x2-x1)=(y2-y1)*(x-x1)
// 这里没有用哈希表缓存，也没法提取一个有效的键来做缓存
func bestLine2(points [][]int) []int {
	n := len(points)
	var ans [2]int
	max := 0
	for i := 0; i < n; i++ {
		if n-i < max {
			break
		}
		for j := i + 1; j < n; j++ {
			if n-j+1 < max {
				break
			}
			xDiff1 := points[j][0] - points[i][0]
			yDiff1 := points[j][1] - points[i][1]
			count := 2
			for k := j + 1; k < n; k++ {
				xDiff2 := points[k][0] - points[i][0]
				yDiff2 := points[k][1] - points[i][1]
				if int64(xDiff1)*int64(yDiff2) == int64(xDiff2)*int64(yDiff1) {
					count++
				}
			}
			if count > max {
				ans[0] = i
				ans[1] = j
				max = count
			}
		}
	}
	return ans[:2]
}
