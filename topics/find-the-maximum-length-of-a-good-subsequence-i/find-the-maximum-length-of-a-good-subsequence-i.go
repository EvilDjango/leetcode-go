// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/9 22:05
package find_the_maximum_length_of_a_good_subsequence_i

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}
	ans := 1
	for i := 0; i < n; i++ {
		dp[i][0] = 1
		for j := 0; j < i; j++ {
			if nums[j] != nums[i] {
				for l := 0; l < k; l++ {
					dp[i][l+1] = max(dp[i][l+1], dp[j][l]+1)
					ans = max(ans, dp[i][l+1])
				}
			} else {
				for l := 0; l <= k; l++ {
					dp[i][l] = max(dp[i][l], dp[j][l]+1)
					ans = max(ans, dp[i][l])
				}
			}
		}
	}
	return ans
}
