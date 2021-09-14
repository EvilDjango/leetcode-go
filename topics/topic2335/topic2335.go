// 面试题 17.19. 消失的两个数字
//给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？
//
//以任意顺序返回这两个数字均可。
//
//示例 1:
//
//输入: [1]
//输出: [2,3]
//示例 2:
//
//输入: [2,3]
//输出: [1,4]
//提示：
//
//nums.length <= 30000
//通过次数8,437提交次数14,504
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/14/21 1:23 PM
package topic2335

// 利用异或
func missingTwo(nums []int) []int {
	n := len(nums)
	aOrb := 0
	for i, num := range nums {
		aOrb ^= i + 1
		aOrb ^= num
	}
	aOrb ^= n + 1
	aOrb ^= n + 2
	mask := aOrb & -aOrb
	a, b := 0, 0
	for i, v := range nums {
		if mask&v > 0 {
			a ^= v
		} else {
			b ^= v
		}
		if mask&(i+1) > 0 {
			a ^= i + 1
		} else {
			b ^= i + 1
		}
	}
	for _, num := range []int{n + 1, n + 2} {
		if mask&num > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

// 先得到两数之和，然后将这组数字拆分为两段，分别求解
func missingTwo2(nums []int) []int {
	n := len(nums) + 2
	sum := n * (n + 1) >> 1
	for _, num := range nums {
		sum -= num
	}
	mid := sum / 2
	sum2 := mid * (mid + 1) >> 1
	for _, num := range nums {
		if num <= mid {
			sum2 -= num
		}
	}
	return []int{sum2, sum - sum2}
}
