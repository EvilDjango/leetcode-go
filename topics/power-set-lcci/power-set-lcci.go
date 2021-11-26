// 面试题 08.04. 幂集
//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
// 输入： nums = [1,2,3]
// 输出：
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//通过次数21,939提交次数26,740
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/26/21 8:50 PM
package power_set_lcci

import "sort"

func subsets(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{{}}
	for i := 0; i < n; i++ {
		size := len(res)
		for j := 0; j < size; j++ {
			length := len(res[j])
			subset := make([]int, length+1)
			copy(subset, res[j])
			subset[length] = nums[i]
			res = append(res, subset)
		}
	}
	return res
}
