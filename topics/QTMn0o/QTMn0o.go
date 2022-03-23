// 剑指 Offer II 010. 和为 k 的子数组
//给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
//
//
//
//示例 1 :
//
//输入:nums = [1,1,1], k = 2
//输出: 2
//解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
//示例 2 :
//
//输入:nums = [1,2,3], k = 3
//输出: 2
//
//
//提示:
//
//1 <= nums.length <= 2 * 104
//-1000 <= nums[i] <= 1000
//-107 <= k <= 107
//
//
//
//注意：本题与主站 560 题相同： https://leetcode-cn.com/problems/subarray-sum-equals-k/
//
//通过次数18,728提交次数42,446
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/23 下午1:22
package QTMn0o

// 前缀和+哈希表
func subarraySum(nums []int, k int) int {
	counts := map[int]int{0: 1}
	prefix := 0
	ans := 0
	for _, num := range nums {
		prefix += num
		if count, ok := counts[prefix-k]; ok {
			ans += count
		}
		counts[prefix]++
	}
	return ans
}
