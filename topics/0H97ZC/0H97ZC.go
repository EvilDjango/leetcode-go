// 剑指 Offer II 075. 数组相对排序
//给定两个数组，arr1 和 arr2，
//
//arr2 中的元素各不相同
//arr2 中的每个元素都出现在 arr1 中
//对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
//
//
//
//示例：
//
//输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//输出：[2,2,2,1,4,3,3,9,6,7,19]
//
//
//提示：
//
//1 <= arr1.length, arr2.length <= 1000
//0 <= arr1[i], arr2[i] <= 1000
//arr2 中的元素 arr2[i] 各不相同
//arr2 中的每个元素 arr2[i] 都出现在 arr1 中
//
//
//注意：本题与主站 1122 题相同：https://leetcode-cn.com/problems/relative-sort-array/
//
//通过次数4,772提交次数6,676
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/27/21 11:03 PM
package _H97ZC

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := map[int]int{}
	for _, num := range arr2 {
		dict[num] = 0
	}
	n := len(arr1)
	i := n - 1
	for j := 0; j <= i; j++ {
		num := arr1[j]
		if _, ok := dict[num]; ok {
			dict[num]++
		} else {
			arr1[i], arr1[j] = arr1[j], arr1[i]
			for k := i; k+1 < n && arr1[k] > arr1[k+1]; k++ {
				arr1[k], arr1[k+1] = arr1[k+1], arr1[k]
			}
			i--
			j--
		}
	}
	i = 0
	for _, num := range arr2 {
		for j := 0; j < dict[num]; j++ {
			arr1[i] = num
			i++
		}
	}
	return arr1
}
