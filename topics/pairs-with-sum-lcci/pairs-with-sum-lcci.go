// 面试题 16.24. 数对和
//设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。
//
//示例 1:
//
//输入: nums = [5,6,5], target = 11
//输出: [[5,6]]
//示例 2:
//
//输入: nums = [5,6,5,6], target = 11
//输出: [[5,6],[5,6]]
//提示：
//
//nums.length <= 100000
//通过次数10,630提交次数22,511
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/22/21 6:53 PM
package pairs_with_sum_lcci

import "sort"

// 利用哈希表
func pairSums(nums []int, target int) [][]int {
	var ans [][]int
	complements := make(map[int]int)
	for _, num := range nums {
		if _, ok := complements[num]; ok {
			ans = append(ans, []int{num, target - num})
			if complements[num] == 1 {
				delete(complements, num)
			} else {
				complements[num]--
			}
		} else {
			complements[target-num]++
		}
	}
	return ans
}

//双指针法，要先排序
func pairSums2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	l, r := 0, len(nums)-1
	for l < r {
		v1, v2 := nums[l], nums[r]
		if v1+v2 > target {
			r--
		} else if v1+v2 < target {
			l++
		} else {
			ans = append(ans, []int{v1, v2})
			l++
			r--
		}
	}
	return ans
}
