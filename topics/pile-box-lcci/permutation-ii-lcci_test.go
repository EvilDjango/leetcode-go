// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/25/21 3:12 PM
package pile_box_lcci

import (
	"leetcode-go/common"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{"qqe"},
			[]string{"eqq", "qeq", "qqe"},
		},
		{
			"",
			args{"ab"},
			[]string{"ab", "ba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.S); !common.EqualIgnoreOrder(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
