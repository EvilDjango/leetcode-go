// 剑指 Offer II 082. 含有重复元素集合的组合
//给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
//
//
//
//示例 1:
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//示例 2:
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//提示:
//
//1 <= candidates.length <= 100
//1 <= candidates[i] <= 50
//1 <= target <= 30
//
//
//注意：本题与主站 40 题相同： https://leetcode-cn.com/problems/combination-sum-ii/
//
//通过次数4,954提交次数7,643
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 1:39 PM
package _sjJUc

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	var ans [][]int
	var dfs func(int, int, []int)
	dfs = func(i, remain int, path []int) {
		if remain == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := i; j < n; j++ {
			num := candidates[j]
			if remain-num < 0 || (j > i && num == candidates[j-1]) {
				continue
			}
			dfs(j+1, remain-num, append(path, num))
		}
	}
	dfs(0, target, nil)
	return ans
}
