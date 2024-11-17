// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/21 23:05
package range_sum_query_mutable

import (
	"testing"
)

func TestNumArray_Update(t *testing.T) {
	obj := Constructor([]int{0, 9, 5, 7, 3})
	r1 := obj.SumRange(4, 4)
	if r1 != 3 {
		t.Error("TestNumArray_Update error")
	}
	r2 := obj.SumRange(2, 4)
	if r2 != 15 {
		t.Error("TestNumArray_Update error")
	}
	r3 := obj.SumRange(3, 3)
	if r3 != 7 {
		t.Error("TestNumArray_Update error")
	}
	obj.Update(4, 5)
	obj.Update(1, 7)
	obj.Update(0, 8)

}
