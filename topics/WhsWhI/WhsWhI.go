//剑指 Offer II 119. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//
//
//示例 1：
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//提示：
//
//0 <= nums.length <= 104
//-109 <= nums[i] <= 109
//
//
//进阶：可以设计并实现时间复杂度为 O(n) 的解决方案吗？
//
//
//
//注意：本题与主站 128 题相同： https://leetcode-cn.com/problems/longest-consecutive-sequence/
//
//通过次数4,602提交次数9,598
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/16/21 11:34 AM
package WhsWhI

func longestConsecutive(nums []int) int {
	// up[i]表示以i结尾的递增序列的长度
	up := map[int]int{}
	// down[i]表示以i结尾的递减序列的长度
	down := map[int]int{}
	ans := 0
	for _, num := range nums {
		if _, ok := up[num]; ok {
			continue
		}
		up[num] = up[num-1] + 1
		down[num] = down[num+1] + 1
		left, right := num-up[num]+1, num+down[num]-1
		length := up[num] + down[num] - 1
		up[right] = length
		down[left] = length
		if length > ans {
			ans = length
		}
	}
	return ans
}

func longestConsecutive2(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	maxLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			length := 1
			for numSet[num+1] {
				length++
				num++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}
