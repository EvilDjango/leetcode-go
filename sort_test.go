// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 12:26 AM
package leetcode_go

import (
	"leetcode-go/common"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestQuickSort(t *testing.T) {
	function := QuickSort
	testQuickSort(t, function)
}

func TestQuickSort2(t *testing.T) {
	function := QuickSort2
	testQuickSort(t, function)
}

func testQuickSort(t *testing.T, function func(arr []int)) {
	for i := 1; i <= 100; i++ {
		ints := common.RandomInts()
		copied := make([]int, len(ints))
		copy(copied, ints)
		sort.Ints(copied)

		function(ints)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(ints, copied) {
				t.Errorf("\ngot:  %v\nwant: %v", ints, copied)
			}
		})
	}
}
