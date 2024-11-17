// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/10/12 22:41
package _1

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = get(num)
	}
	return ans
}

func get(num int) int {
	if num == 2 {
		return -1
	}
	i := 0
	for ; i|(i+1) != num; i++ {
	}
	return i
}
