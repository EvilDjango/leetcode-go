// 剑指 Offer II 069. 山峰数组的顶部
//符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：
//
//arr.length >= 3
//存在 i（0 < i < arr.length - 1）使得：
//arr[0] < arr[1] < ... arr[i-1] < arr[i]
//arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//给定由整数组成的山峰数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i ，即山峰顶部。
//
//
//
//示例 1：
//
//输入：arr = [0,1,0]
//输出：1
//示例 2：
//
//输入：arr = [1,3,5,4,2]
//输出：2
//示例 3：
//
//输入：arr = [0,10,5,2]
//输出：1
//示例 4：
//
//输入：arr = [3,4,5,1]
//输出：2
//示例 5：
//
//输入：arr = [24,69,100,99,79,78,67,36,26,19]
//输出：2
//
//
//提示：
//
//3 <= arr.length <= 104
//0 <= arr[i] <= 106
//题目数据保证 arr 是一个山脉数组
//
//
//进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？
//
//
//
//注意：本题与主站 852 题相同：https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/
//
//通过次数38,679提交次数53,635
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/28/21 10:59 PM
package B1IidL

func peakIndexInMountainArray(arr []int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := (high-low)>>1 + low
		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
