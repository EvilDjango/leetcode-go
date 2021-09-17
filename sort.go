// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 12:10 AM
package leetcode_go

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	l, r, anchor := left, right, arr[left]
	for l < r {
		for l < r && arr[r] >= anchor {
			r--
		}
		for l < r && arr[l] <= anchor {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[l], arr[left] = arr[left], arr[l]
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
}

// 另一种写法
func QuickSort2(arr []int) {
	quickSort2(arr, 0, len(arr)-1)
}

func quickSort2(arr []int, left, right int) {
	if left >= right {
		return
	}
	l, r, anchor := left, right, arr[left]
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
	quickSort2(arr, left, l-1)
	quickSort2(arr, l+1, right)
}
