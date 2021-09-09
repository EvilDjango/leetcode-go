package topic301

import (
	"reflect"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/24/21 2:47 PM
*/

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{"()())()"},
			[]string{"(())()", "()()()"},
		},
		{
			"",
			args{")("},
			[]string{""},
		},
		{
			"",
			args{"(a)())()"},
			[]string{"(a())()", "(a)()()"},
		},
		{
			"",
			args{")()("},
			[]string{"()"},
		},
		{
			"",
			args{"(()("},
			[]string{"()"},
		},
		{
			"",
			args{"(((k()(("},
			[]string{"k()", "(k)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
