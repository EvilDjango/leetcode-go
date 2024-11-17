// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/30 22:32
package sum_of_digit_differences_of_all_pairs

func sumDigitDifferences(nums []int) int64 {
	var ans int64 = 0
	count := make([]int64, 10)
	for base := 1; nums[0] >= base; base = base * 10 {
		for _, num := range nums {
			count[num/base%10]++
		}
		for j := 0; j < 10; j++ {
			ans += count[j] * (int64(len(nums)) - count[j])
		}
		clear(count)
	}
	return ans / 2
}
