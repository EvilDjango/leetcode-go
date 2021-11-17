// 面试题 08.13. 堆箱子
//堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。
//
//输入使用数组[wi, di, hi]表示每个箱子。
//
//示例1:
//
// 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
// 输出：6
//示例2:
//
// 输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
// 输出：10
//提示:
//
//箱子的数目不大于3000个。
//通过次数7,160提交次数14,266
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/16/21 8:54 AM
package pile_box_lcci

import "sort"

// 动态规划
func pileBox(box [][]int) int {
	sort.Slice(box, func(i, j int) bool {
		if box[i][0] == box[j][0] {
			return box[i][1] > box[j][1]
		}
		return box[i][0] < box[j][0]
	})
	n := len(box)
	dp := make([]int, n)
	maxH := 0
	for i := 0; i < n; i++ {
		dp[i] = box[i][2]
		for j := 0; j < i; j++ {
			if box[j][1] < box[i][1] && box[j][2] < box[i][2] {
				dp[i] = max(dp[i], dp[j]+box[i][2])
			}
		}
		maxH = max(maxH, dp[i])
	}
	return maxH
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
