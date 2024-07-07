// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/6 上午11:27
package count_alternating_subarrays

func countAlternatingSubarrays(nums []int) int64 {
	res := 0
	prev := -1
	length := 0
	for _, num := range nums {
		if num != prev {
			length++
		} else {
			length = 1
		}
		prev = num
		res += length
	}
	return int64(res)
}
