// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/24/21 10:29 PM
package bracket_lcci

import (
	"fmt"
	"leetcode-go/common"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{4},
			[]string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
		{
			"",
			args{3},
			[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(generateParenthesis(tt.args.n))

			if got := generateParenthesis(tt.args.n); !common.EqualIgnoreOrder(tt.want, got) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
