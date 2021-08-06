package leetcode_go

import (
	"leetcode-go/common"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/29/21 5:38 PM
*/

func Test_generatePalindromes(t *testing.T) {
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
			args{"aabb"},
			[]string{"abba", "baab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePalindromes(tt.args.s); !common.EqualIgnoreOrder(got, tt.want) {
				t.Errorf("generatePalindromes() = %v, want %v", got, tt.want)
			}
		})
	}
}
