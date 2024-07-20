// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/20 下午8:29
package permutations

func permute(nums []int) [][]int {
	var ans [][]int
	var backtrack func(int)
	n := len(nums)
	permutation := make([]int, n)
	used := make([]bool, n)
	backtrack = func(index int) {
		if index == n {
			copied := make([]int, n)
			copy(copied, permutation)
			ans = append(ans, copied)
			return
		}
		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			permutation[index] = num
			backtrack(index + 1)
			used[i] = false
		}
	}
	backtrack(0)
	return ans
}

func permute2(nums []int) [][]int {
	var ans [][]int
	var backtrack func(int)
	n := len(nums)
	backtrack = func(index int) {
		if index == n {
			copied := make([]int, n)
			copy(copied, nums)
			ans = append(ans, copied)
			return
		}
		for i := index; i < n; i++ {
			nums[index], nums[i] = nums[i], nums[index]
			backtrack(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	backtrack(0)
	return ans
}
