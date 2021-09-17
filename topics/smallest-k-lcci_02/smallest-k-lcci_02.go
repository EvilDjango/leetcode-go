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
	quickSort(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func quickSort(arr []int, k, left, right int) {
	if left >= right {
		return
	}
	rnd := left + rand.Intn(right-left+1)
	arr[left], arr[rnd] = arr[rnd], arr[left]

	l, r, anchor := left, right, arr[left]
	for l < r {
		for l < r && arr[r] > anchor {
			r--
		}
		for l < r && arr[l] <= anchor {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[left], arr[l] = arr[l], arr[left]
	if l > k-1 {
		quickSort(arr, k, left, l-1)
	} else if l < k {
		quickSort(arr, k, l+1, right)
	}
}

// 另一种快排写法
func quickSort2(arr []int, k, left, right int) {
	l, r := left, right
	if l >= r {
		return
	}
	rnd := left + rand.Intn(right-left+1)
	arr[left], arr[rnd] = arr[rnd], arr[left]
	anchor := arr[l]
	for l < r {
		for l < r && arr[r] >= anchor {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= anchor {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = anchor
	if l < k-1 {
		quickSort2(arr, k, l+1, right)
	} else if l > k {
		quickSort2(arr, k, left, l-1)
	}
}
