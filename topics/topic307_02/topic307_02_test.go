// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/3/21 11:53 AM
package topic307_02

import (
	"fmt"
	"testing"
)

func TestNumArray(t *testing.T) {
	na := Constructor([]int{0, 1, 2, 3, 4, 5})
	testCases := [][]int{{0, 5}, {2, 3}, {2, 4}}
	wants := []int{15, 5, 9}
	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		t.Run(fmt.Sprintf("%d-%d", testCase[0], testCase[1]), func(t *testing.T) {
			res := na.SumRange(testCase[0], testCase[1])
			if res != wants[i] {
				t.Errorf("got %d, want %d", res, wants[i])
			}
		})
	}
	na.Update(0, 1)

	testCases = [][]int{{0, 5}, {2, 3}, {2, 4}}
	wants = []int{16, 5, 9}
	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		t.Run(fmt.Sprintf("%d-%d", testCase[0], testCase[1]), func(t *testing.T) {
			res := na.SumRange(testCase[0], testCase[1])
			if res != wants[i] {
				t.Errorf("got %d, want %d", res, wants[i])
			}
		})
	}

}
