package heap

import (
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/19/21 8:13 PM
*/

func TestMinHeap(t *testing.T) {
	minHeap := Init([]int{5, 6, 7, 4, 2, 3}, false)
	minHeap.Push(1)
	t.Run("", func(t *testing.T) {
		size := minHeap.Len()
		for i := 1; i <= size; i++ {
			v := minHeap.Pop()
			if i != v {
				t.Errorf("get %d, want %d", v, i)
			}
		}
	})
}

func TestMaxHeap(t *testing.T) {
	minHeap := Init([]int{5, 6, 7, 4, 1, 2, 3}, true)
	t.Run("", func(t *testing.T) {
		size := minHeap.Len()
		for i := size; i > 0; i-- {
			v := minHeap.Pop()
			if i != v {
				t.Errorf("get %d, want %d", v, i)
			}
		}
	})
}
