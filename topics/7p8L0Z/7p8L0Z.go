// 剑指 Offer II 084. 含有重复元素集合的全排列
//给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
//
//
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//示例 2：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10
//
//
//注意：本题与主站 47 题相同： https://leetcode-cn.com/problems/permutations-ii/
//
//通过次数4,695提交次数6,846
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/23/21 11:35 PM
package _p8L0Z

import "sort"

// 利用字典去重
func permuteUnique(nums []int) [][]int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}
	var ans [][]int
	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		if i == len(nums) {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for num, count := range dict {
			if count == 0 {
				continue
			}
			dict[num]--
			dfs(i+1, append(path, num))
			dict[num]++
		}
	}
	dfs(0, nil)
	return ans
}

// 参考官方题解
func permuteUnique2(nums []int) [][]int {
	n := len(nums)
	visited := make([]bool, n)
	sort.Ints(nums)
	var ans [][]int
	var dfs func(int, []int)
	dfs = func(count int, path []int) {
		if count == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j := 0; j < n; j++ {
			if visited[j] || (j > 0 && nums[j] == nums[j-1] && !visited[j-1]) {
				continue
			}
			visited[j] = true
			dfs(count+1, append(path, nums[j]))
			visited[j] = false
		}
	}
	dfs(0, nil)
	return ans
}
