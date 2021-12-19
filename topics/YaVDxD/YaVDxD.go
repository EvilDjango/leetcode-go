// 剑指 Offer II 102. 加减的目标值
//给定一个正整数数组 nums 和一个整数 target 。
//
//向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
//
//
//示例 1：
//
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
//示例 2：
//
//输入：nums = [1], target = 1
//输出：1
//
//
//提示：
//
//1 <= nums.length <= 20
//0 <= nums[i] <= 1000
//0 <= sum(nums[i]) <= 1000
//-1000 <= target <= 1000
//
//
//注意：本题与主站 494 题相同： https://leetcode-cn.com/problems/target-sum/
//
//通过次数2,812提交次数4,736
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/18/21 9:04 PM
package YaVDxD

// 动态规划
// 可以认为是将nums分组，前面为+的一组总和记为a，前面为-的一组总和记为b，记nums总和为sum，那么有a+b=sum,a-b=target
// 计算可得a=(sum+target)/2 b=(sum-target)/2
// 利用这个关系，我们可以仅关注a或者b，减少运算
// sum+target==2a，必为偶数
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum-target)%2 != 0 {
		return 0
	}
	// dp[i]表示挑选一部分元素凑成总和的方案数
	dp := make([]int, (sum+target)/2+1)
	dp[0] = 1
	for _, num := range nums {
		for i := len(dp) - 1; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[len(dp)-1]
}
