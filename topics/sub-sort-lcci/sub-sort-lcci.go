// 面试题 16.16. 部分排序
//给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。
//
//示例：
//
//输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
//输出： [3,9]
//提示：
//
//0 <= len(array) <= 1000000
//通过次数14,699提交次数32,421
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/25/21 6:32 PM
package sub_sort_lcci

import "math"

// 记满足条件的区间为[l,r]，那么可以推出，l是左起第一个满足array[i]大于i右侧元素最小值的下标，r是右起第一个满足array[i]小于i左侧元素最大值的下标
func subSort(array []int) []int {
	n := len(array)
	l, r := -1, -1
	min := math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		if array[i] > min {
			l = i
		} else {
			min = array[i]
		}
	}
	max := math.MinInt32
	for i, num := range array {
		if num < max {
			r = i
		} else {
			max = num
		}
	}
	return []int{l, r}
}
