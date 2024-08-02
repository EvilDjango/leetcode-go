// Oops, forgot to write comments. Good luck, bro. 
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/31 下午10:13
package minimum_rectangles_to_cover_points

import (
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	ans := 0
	bound := -1
	for _, p := range points {
		if bound < p[0] {
			ans++
			bound = p[0] + w
		}
	}
	return ans
}
