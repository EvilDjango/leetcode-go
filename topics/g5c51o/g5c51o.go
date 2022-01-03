// 剑指 Offer II 060. 出现频率最高的 k 个数字
//给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。
//
//
//
//示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//示例 2:
//
//输入: nums = [1], k = 1
//输出: [1]
//
//
//提示：
//
//1 <= nums.length <= 105
//k 的取值范围是 [1, 数组中不相同的元素的个数]
//题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//进阶：所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
//
//
//
//注意：本题与主站 347 题相同：https://leetcode-cn.com/problems/top-k-frequent-elements/
//
//通过次数4,502提交次数6,601
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/31/21 11:16 AM
package g5c51o

import (
	"container/heap"
	"math/rand"
	"sort"
)

type hp [][]int

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	a := *h
	return a[i][1] < a[j][1]
}
func (h *hp) Swap(i, j int) {
	a := *h
	a[i], a[j] = a[j], a[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.([]int))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

// 排序， 时间复杂度为NlgN，不符合要求
func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	pairs := make([][]int, 0, len(count))
	for num, n := range count {
		pairs = append(pairs, []int{num, n})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})
	ans := make([]int, 0, k)
	for _, pair := range pairs {
		ans = append(ans, pair[0])
		if len(ans) == k {
			break
		}
	}
	return ans
}

// 小顶堆
func topKFrequent2(nums []int, k int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	h := &hp{}
	heap.Init(h)
	for num, n := range count {
		if h.Len() < k {
			heap.Push(h, []int{num, n})
		} else if n > (*h)[0][1] {
			heap.Pop(h)
			heap.Push(h, []int{num, n})
		}
	}
	ans := make([]int, 0, k)
	for _, pair := range *h {
		ans = append(ans, pair[0])
	}
	return ans
}

// 利用快速排序的思想
func topKFrequent3(nums []int, k int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	pairs := make([][]int, 0, len(count))
	for num, n := range count {
		pairs = append(pairs, []int{num, n})
	}
	n := len(pairs)
	if k < n {
		l, r := 0, n-1
		for l < r {
			left, right := l, r
			i := rand.Intn(r-l) + l
			swap(pairs, i, r)
			base := pairs[r]
			for l < r {
				for l < r && pairs[l][1] >= base[1] {
					l++
				}
				swap(pairs, l, r)
				for l < r && pairs[r][1] <= base[1] {
					r--
				}
				swap(pairs, l, r)
			}
			pairs[r] = base
			if l < k-1 {
				l, r = l+1, right
			} else if l > k {
				l, r = left, l-1
			} else {
				break
			}
		}
	} else {
		k = n
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = pairs[i][0]
	}
	return ans
}

func swap(arr [][]int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
