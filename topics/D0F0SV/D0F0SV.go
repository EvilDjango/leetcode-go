// 剑指 Offer II 104. 排列的数目
//给定一个由 不同 正整数组成的数组 nums ，和一个目标整数 target 。请从 nums 中找出并返回总和为 target 的元素组合的个数。数组中的数字可以在一次排列中出现任意次，但是顺序不同的序列被视作不同的组合。
//
//题目数据保证答案符合 32 位整数范围。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3], target = 4
//输出：7
//解释：
//所有可能的组合为：
//(1, 1, 1, 1)
//(1, 1, 2)
//(1, 2, 1)
//(1, 3)
//(2, 1, 1)
//(2, 2)
//(3, 1)
//请注意，顺序不同的序列被视作不同的组合。
//示例 2：
//
//输入：nums = [9], target = 3
//输出：0
//
//
//提示：
//
//1 <= nums.length <= 200
//1 <= nums[i] <= 1000
//nums 中的所有元素 互不相同
//1 <= target <= 1000
//
//
//进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
//
//
//
//注意：本题与主站 377 题相同：https://leetcode-cn.com/problems/combination-sum-iv/
//
//通过次数2,139提交次数3,634
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/17/21 2:26 PM
package D0F0SV

import "sort"

// 记忆化递归
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	cache := map[int]int{}
	var dfs func(int) int
	dfs = func(remain int) int {
		if remain == 0 {
			return 1
		}
		if v, ok := cache[remain]; ok {
			return v
		}
		count := 0
		for i := 0; i < len(nums) && nums[i] <= remain; i++ {
			count += dfs(remain - nums[i])
		}
		cache[remain] = count
		return count
	}
	return dfs(target)
}

// 动态规划
func combinationSum42(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	sort.Ints(nums)
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums) && nums[j] <= i; j++ {
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}
