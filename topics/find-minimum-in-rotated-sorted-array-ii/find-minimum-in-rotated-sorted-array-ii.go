// 154. 寻找旋转排序数组中的最小值 II
//已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
//若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
//若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
//给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
//
//
//
//示例 1：
//
//输入：nums = [1,3,5]
//输出：1
//示例 2：
//
//输入：nums = [2,2,2,0,1]
//输出：0
//
//
//提示：
//
//n == nums.length
//1 <= n <= 5000
//-5000 <= nums[i] <= 5000
//nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
//
//
//进阶：
//
//这道题是 寻找旋转排序数组中的最小值 的延伸题目。
//允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
//通过次数117,578提交次数221,206
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/14/21 3:26 PM
package find_minimum_in_rotated_sorted_array_ii

// 元素重复带来的最大的问题是，分割点前后的元素可能相同，这会导致我们无法判断某一个位置处于左边还是右边。那么我们只需要在这种情况下把某一边的重复元素跳过即可
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] == nums[r] {
		for ; l < r && nums[l] == nums[l+1]; l++ {
			// empty body
		}
	}
	for l < r {
		mid := (r-l)>>1 + l
		// mid在左侧
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

// 参考官方题解
func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		pivot := (r-l)>>1 + l
		// pivot在左侧
		if nums[pivot] > nums[r] {
			l = pivot + 1
		} else if nums[pivot] < nums[r] {
			r = pivot
		} else {
			r--
		}
	}
	return nums[l]
}
