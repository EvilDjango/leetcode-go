// 剑指 Offer II 101. 分割等和子集
//给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。
//
//
//
//示例 1：
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：nums 可以分割成 [1, 5, 5] 和 [11] 。
//示例 2：
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：nums 不可以分为和相等的两部分
//
//
//提示：
//
//1 <= nums.length <= 200
//1 <= nums[i] <= 100
//
//
//注意：本题与主站 416 题相同： https://leetcode-cn.com/problems/partition-equal-subset-sum/
//
//通过次数3,799提交次数7,778
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/18/21 9:05 PM
package NUPfPr

// 记忆化递归
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, sum/2+1)
	}
	var dfs func(int, int) bool
	// 从index开始是否可以选取一部分元素使其和等于remain
	dfs = func(index int, remain int) bool {
		if index == n {
			return remain == 0
		}
		if remain <= 0 {
			return false
		}
		if cache[index][remain] == 0 {
			ok := dfs(index+1, remain) || dfs(index+1, remain-nums[index])
			if ok {
				cache[index][remain] = 1
			} else {
				cache[index][remain] = 2
			}
		}
		return cache[index][remain] == 1
	}
	return dfs(0, sum/2)
}

// 动态规划
func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	dp := make([]bool, sum/2+1)
	dp[0] = true
	for _, num := range nums {
		for j := num; j <= sum/2; j++ {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[sum/2]
}
