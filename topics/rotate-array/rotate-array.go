// 189. 轮转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
//
//
//示例 1:
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]
//
//
//提示：
//
//1 <= nums.length <= 105
//-231 <= nums[i] <= 231 - 1
//0 <= k <= 105
//
//
//进阶：
//
//尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
//你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//通过次数370,051提交次数825,444
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/14/21 11:55 AM
package rotate_array

// 翻转
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if n == 0 || k == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 环状替换
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	times := gcd(n, k)
	for i := 0; i < times; i++ {
		temp := nums[i]
		for j := (i + k) % n; j != i; j = (j + k) % n {
			temp, nums[j] = nums[j], temp
		}
		nums[i] = temp
	}
}

func gcd(i, j int) int {
	for i%j != 0 {
		i, j = j, i%j
	}
	return j
}
