// 面试题 16.06. 最小差
//给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差
//
//
//
//示例：
//
//输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
//输出：3，即数值对(11, 8)
//
//
//提示：
//
//1 <= a.length, b.length <= 100000
//-2147483648 <= a[i], b[i] <= 2147483647
//正确结果在区间 [0, 2147483647] 内
//通过次数12,874提交次数30,360
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/15/21 8:39 AM
package smallest_difference_lcci

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	i, j, diff := 0, 0, 0
	ans := math.MaxInt32
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			diff = b[j] - a[i]
			i++
		} else {
			diff = a[i] - b[j]
			j++
		}
		if diff == 0 {
			return 0
		}
		if diff < ans {
			ans = diff
		}
	}
	return ans
}
