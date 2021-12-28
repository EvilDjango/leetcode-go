// 剑指 Offer II 070. 排序数组中只出现一次的数字
//给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。
//
//
//
//示例 1:
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//示例 2:
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
//
//
//
//
//提示:
//
//1 <= nums.length <= 105
//0 <= nums[i] <= 105
//
//
//进阶: 采用的方案可以在 O(log n) 时间复杂度和 O(1) 空间复杂度中运行吗？
//
//
//
//注意：本题与主站 540 题相同：https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
//
//通过次数5,069提交次数7,789
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/28/21 10:03 PM
package skFtm2

// 异或
func singleNonDuplicate(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

// 二分搜索
func singleNonDuplicate2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (high-low)>>1 + low
		mid /= 2
		mid *= 2
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}
