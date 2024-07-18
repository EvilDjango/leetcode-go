// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/10 下午10:12
package count_the_number_of_incremovable_subarrays_i

func incremovableSubarrayCount(nums []int) int {
	n := len(nums)
	l := 1
	for l < n && nums[l] > nums[l-1] {
		l++
	}
	if l == n {
		return n * (n + 1) / 2
	}
	res := l + 1
	r := 0
	for r < n {
		if r > 1 && nums[n-r] >= nums[n-r+1] {
			break
		}
		for l > 0 && nums[l-1] >= nums[n-r] {
			l--
		}
		res += l + 1
		r++
	}
	return res
}
