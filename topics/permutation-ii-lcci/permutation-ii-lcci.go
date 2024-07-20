// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/20 下午8:55
package permutation_ii_lcci

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	var dfs func(int)
	used := make([]bool, n)
	perm := make([]int, n)
	dfs = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, perm...))
			return
		}
		for i, v := range nums {
			if used[i] || (i > 0 && v == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			perm[index] = v
			dfs(index + 1)
			used[i] = false
		}
	}

	sort.Ints(nums)
	dfs(0)
	return ans
}

func permuteUnique2(nums []int) [][]int {
	var ans [][]int
	var dfs func(int)
	dfs = func(index int) {
		if index == len(nums) {
			copied := make([]int, len(nums))
			copy(copied, nums)
			ans = append(ans, copied)
			return
		}
		used := make(map[int]bool)

		for i := index; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			nums[i], nums[index] = nums[index], nums[i]
			dfs(index + 1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	dfs(0)
	return ans
}
