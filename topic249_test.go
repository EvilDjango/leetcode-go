package leetcode_go

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/22/21 8:46 PM
*/

func Test_canTransform(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"a",
			args{"a", "n"},
			true,
		},
		{
			"ab",
			args{"ab", "bc"},
			true,
		},
		{
			"eqdf",
			args{"eqdf", "qcpr"},
			true,
		},
		{
			"rbnzendwnoklpyyyauemm",
			args{"rbnzendwnoklpyyyauemm", "fpbnsbrkbcyzdmmmoisaa"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTransform(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("canTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
