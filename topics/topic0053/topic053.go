// 计算一个数组中的最大区间和
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/10/21 5:07 PM
package topic0053

import "math"

func maxSubArray(nums []int) int {
	maxSum, sum := math.MinInt32, 0
	for _, num := range nums {
		sum = max(num, sum+num)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
