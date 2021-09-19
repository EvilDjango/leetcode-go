// 面试题 17.10. 主要元素
//数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
//
//
//
//示例 1：
//
//输入：[1,2,5,9,5,9,5,5,5]
//输出：5
//示例 2：
//
//输入：[3,2]
//输出：-1
//示例 3：
//
//输入：[2,2,1,1,1,2,2]
//输出：2
//通过次数66,469提交次数116,414
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/19/21 2:29 PM
package find_majority_element_lcci

func majorityElement(nums []int) int {
	ans, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			ans, cnt = num, 1
		} else if ans == num {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt == 0 {
		return -1
	}
	cnt = 0
	for _, num := range nums {
		if num == ans {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt > 0 {
		return ans
	}
	return -1
}
