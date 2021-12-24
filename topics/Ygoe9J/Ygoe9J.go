// 剑指 Offer II 081. 允许重复选择元素的组合
//给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
//
//candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
//
//对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
//
//
//
//示例 1：
//
//输入: candidates = [2,3,6,7], target = 7
//输出: [[7],[2,2,3]]
//示例 2：
//
//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]
//示例 3：
//
//输入: candidates = [2], target = 1
//输出: []
//示例 4：
//
//输入: candidates = [1], target = 1
//输出: [[1]]
//示例 5：
//
//输入: candidates = [1], target = 2
//输出: [[1,1]]
//
//
//提示：
//
//1 <= candidates.length <= 30
//1 <= candidates[i] <= 200
//candidate 中的每个元素都是独一无二的。
//1 <= target <= 500
//
//
//注意：本题与主站 39 题相同： https://leetcode-cn.com/problems/combination-sum/
//
//通过次数6,297提交次数7,998
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 2:19 PM
package Ygoe9J

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var dfs func(int, int, []int)
	dfs = func(i, rest int, nums []int) {
		if rest == 0 {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for j := i; j < len(candidates) && candidates[j] <= rest; j++ {
			dfs(j, rest-candidates[j], append(nums, candidates[j]))
		}
	}
	dfs(0, target, nil)
	return ans
}
