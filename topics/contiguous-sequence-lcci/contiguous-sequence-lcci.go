// 面试题 16.17. 连续数列
//给定一个整数数组，找出总和最大的连续数列，并返回总和。
//
//示例：
//
//输入： [-2,1,-3,4,-1,2,1,-5,4]
//输出： 6
//解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
//进阶：
//
//如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
//
//通过次数34,205提交次数56,933
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/24/21 10:24 AM
package contiguous_sequence_lcci

import "math"

// 动态规划 dp[i]表示以元素i结尾的最大字段和，那么 dp[i+1】=max(dp[i]+nums[i],nums[i])
// 由于dp[i]只跟前一个状态有关，所以可以只记录前一状态节省空间
func maxSubArray(nums []int) int {
	max := math.MinInt32
	sum := 0
	for _, num := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += num
		if sum > max {
			max = sum
		}
	}
	return max
}

// 分治法
func maxSubArray2(nums []int) int {
	return getStats(nums, 0, len(nums)-1).mSum
}

func getStats(nums []int, l, r int) *Stats {
	if l == r {
		return &Stats{nums[l], nums[l], nums[l], nums[l]}
	}
	mid := (r-l)/2 + l
	lStats, rStats := getStats(nums, l, mid), getStats(nums, mid+1, r)
	lSum := max(lStats.lSum, lStats.sum+rStats.lSum)
	rSum := max(rStats.rSum, lStats.rSum+rStats.sum)
	sum := lStats.sum + rStats.sum
	mSum := max(max(lStats.mSum, rStats.mSum), lStats.rSum+rStats.lSum)
	return &Stats{lSum, rSum, sum, mSum}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type Stats struct {
	// 分别表示从左端起的最大子段和，截至右端的最大子段和，序列总和，最大子段和
	lSum, rSum, sum, mSum int
}
