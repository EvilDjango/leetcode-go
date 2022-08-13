// 768. 最多能完成排序的块 II
//这个问题和“最多能完成排序的块”相似，但给定数组中的元素可以重复，输入数组最大长度为2000，其中的元素最大为10**8。
//
//arr是一个可能包含重复元素的整数数组，我们将这个数组分割成几个“块”，并将这些块分别进行排序。之后再连接起来，使得连接的结果和按升序排序后的原数组相同。
//
//我们最多能将数组分成多少块？
//
//示例 1:
//
//输入: arr = [5,4,3,2,1]
//输出: 1
//解释:
//将数组分成2块或者更多块，都无法得到所需的结果。
//例如，分成 [5, 4], [3, 2, 1] 的结果是 [4, 5, 1, 2, 3]，这不是有序的数组。
//示例 2:
//
//输入: arr = [2,1,3,4,4]
//输出: 4
//解释:
//我们可以把它分成两块，例如 [2, 1], [3, 4, 4]。
//然而，分成 [2, 1], [3], [4], [4] 可以得到最多的块数。
//注意:
//
//arr的长度在[1, 2000]之间。
//arr[i]的大小在[0, 10**8]之间。
//通过次数13,290提交次数23,819
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/8/13 上午11:19
package max_chunks_to_make_sorted_ii

import "sort"

// 统计最大值和最大值出现的次数
func maxChunksToSorted(arr []int) int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	ans := 0
	max, curr := [2]int{}, [2]int{}
	for i, v := range arr {
		if i == 0 || sorted[i-1] != sorted[i] {
			curr[0] = sorted[i]
			curr[1] = 1
		} else {
			curr[1]++
		}
		if v > max[0] {
			max[0] = v
			max[1] = 1
		} else if v == max[0] {
			max[1]++
		}
		if max == curr {
			ans++
		}
	}
	return ans
}

// 参考优秀题解，统计前缀和
func maxChunksToSorted2(arr []int) int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	ans := 0
	sumDiff := 0
	for i, v := range arr {
		sumDiff += v - sorted[i]
		if sumDiff == 0 {
			ans++
		}
	}
	return ans
}

// 单调栈
func maxChunksToSorted3(arr []int) int {
	var stack []int
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
		for len(stack) > 0 && stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, max)
	}
	return len(stack)
}
