// 452. 用最少数量的箭引爆气球
//在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。
//
//一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
//
//给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。
//
//
//示例 1：
//
//输入：points = [[10,16],[2,8],[1,6],[7,12]]
//输出：2
//解释：对于该样例，x = 6 可以射爆 [2,8],[1,6] 两个气球，以及 x = 11 射爆另外两个气球
//示例 2：
//
//输入：points = [[1,2],[3,4],[5,6],[7,8]]
//输出：4
//示例 3：
//
//输入：points = [[1,2],[2,3],[3,4],[4,5]]
//输出：2
//示例 4：
//
//输入：points = [[1,2]]
//输出：1
//示例 5：
//
//输入：points = [[2,3],[2,3]]
//输出：1
//
//
//提示：
//
//1 <= points.length <= 104
//points[i].length == 2
//-231 <= xstart < xend <= 231 - 1
//通过次数102,321提交次数201,290
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/16/21 8:09 PM
package minimum_number_of_arrows_to_burst_balloons

import (
	"math"
	"sort"
)

// 贪心算法 先按照左端点排序。然后从左到右依次遍历。从第一个气球开始，记右边界为right，如果下一个气球的左边界小于等于right，那么两个气球可以用同一枪同时打中，
// 继续遍历下一个气球，如果其左边界还是小于等于right，此时三个气球可以一枪同时打中吗？答案是不一定，因为第二和第三个气球不一定有公共部分。
// 怎么解决这个问题呢？在上一步的时候，right要取min(right,第二个气球的右端点），这样第三个气球的左边界如果还是小于等于right，那么三个气球也可以一枪同时打中。
// 如果后面一个气球左边界大于right的情况就更简单了，显然需要另外开一枪，那么这又是一个新的开始了。
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	count := 1
	right := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= right {
			if points[i][1] < right {
				right = points[i][1]
			}
		} else {
			count++
			right = points[i][1]
		}
	}
	return count
}

// 另一种贪心算法 这个算法是求最长的不互相覆盖的序列长度，其实也就等价于最小的开枪数。
// 先按照右端点排序。一开始总是让右端点最小的元素作为序列的开头，因为这样可以给右侧留出最多的空间来容纳更多的气球。
// 确定一个气球之后，用right来记录起右端点，接着从剩余的气球中选取右端点最小的即可，遇到会与前面的气球互相覆盖的直接跳过即可。
func findMinArrowShots2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 0
	right := math.MinInt64
	for _, point := range points {
		if point[0] > right {
			count++
			right = point[1]
		}
	}
	return count
}
