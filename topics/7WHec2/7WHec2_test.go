// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 11:52 PM
package _WHec2

import (
	. "leetcode-go/container/linklist"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	test(t, sortList)
}

func Test_sortList2(t *testing.T) {
	test(t, sortList2)
}

func Test_sortList3(t *testing.T) {
	test(t, sortList3)
}

func test(t *testing.T, f func(head *ListNode) *ListNode) {
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
			args{New(4, 2, 1, 3)},
			New(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
