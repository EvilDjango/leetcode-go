// 面试题 16.21. 交换和
//给定两个整数数组，请交换一对数值（每个数组中取一个数值），使得两个数组所有元素的和相等。
//
//返回一个数组，第一个元素是第一个数组中要交换的元素，第二个元素是第二个数组中要交换的元素。若有多个答案，返回任意一个均可。若无满足条件的数值，返回空数组。
//
//示例:
//
//输入: array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
//输出: [1, 3]
//示例:
//
//输入: array1 = [1, 2, 3], array2 = [4, 5, 6]
//输出: []
//提示：
//
//1 <= array1.length, array2.length <= 100000
//通过次数9,363提交次数20,143
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/23/21 12:48 PM
package sum_swap_lcci

func findSwapValues(array1 []int, array2 []int) []int {
	sum1 := 0
	count1 := map[int]int{}
	for _, num := range array1 {
		sum1 += num
		count1[num]++
	}
	sum2 := 0
	for _, num := range array2 {
		sum2 += num
	}
	diff := sum2 - sum1
	if diff%2 == 1 {
		return nil
	}
	delta := diff >> 1
	for _, num := range array2 {
		pair := num - delta
		if count1[pair] > 0 {
			return []int{pair, num}
		}
	}
	return nil
}
