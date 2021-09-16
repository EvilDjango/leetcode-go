// 面试题 17.14. 最小K个数
//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
//示例：
//
//输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//提示：
//
//0 <= len(arr) <= 100000
//0 <= k <= min(100000, len(arr))
//通过次数75,332提交次数125,988
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/16/21 6:36 PM
package smallest_k_lcci

import "math/rand"

// 利用快排的思想
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	quickSort3(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func quickSort1(arr []int, k, l, r int) {
	i, j := l, r
	if i >= j {
		return
	}

	rnd := l + rand.Intn(r-l+1)
	arr[l], arr[rnd] = arr[rnd], arr[l]

	anchor := arr[i]

	for i < j {
		for i < j && arr[j] >= anchor {
			j--
		}

		arr[i] = arr[j]

		for i < j && arr[i] <= anchor {
			i++
		}

		arr[j] = arr[i]
	}

	arr[i] = anchor

	if i < k-1 {
		quickSort1(arr, k, i+1, r)
	}

	if i > k {
		quickSort1(arr, k, l, i-1)
	}
}

func quickSort3(arr []int, k, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	base := arr[l]
	for l < r {
		for l < r && arr[r] > base {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= base {
			l++
		}
		arr[r] = arr[l]
	}
	arr[r] = base
	if r > k-1 {
		quickSort1(arr, r-1, k, left)
	} else if r < k-1 {
		quickSort1(arr, right, k, r+1)
	}
}

// 另一种快排写法
func quickSort2(arr []int, k, left, right int) {
	if left > right {
		return
	}
	l, r, base := left+1, right, arr[left]
	for l <= r {
		for l <= r && arr[r] > base {
			r--
		}
		for l <= r && arr[l] <= base {
			l++
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[left], arr[r] = arr[r], base
	if r > k-1 {
		quickSort1(arr, r-1, k, left)
	} else if r < k-1 {
		quickSort1(arr, right, k, r+1)
	}
}
