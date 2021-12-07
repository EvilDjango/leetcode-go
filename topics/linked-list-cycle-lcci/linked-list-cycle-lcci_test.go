// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/7/21 4:01 PM
package linked_list_cycle_lcci

import (
	. "leetcode-go/container/linklist"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"",
			args{CreateCircle(1, 0, 1, 2, 3, 4)},
			New(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); got.Val != tt.want.Val {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
