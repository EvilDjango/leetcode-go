package topic0279

import (
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/11/21 12:46 PM
*/

func Test_numSquares(t *testing.T) {
	test279(t, numSquares)
}
func Test_numSquares2(t *testing.T) {
	test279(t, numSquares2)
}

func test279(t *testing.T, function func(n int) int) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{12},
			3,
		},
		{
			"",
			args{13},
			2,
		},
		{
			"",
			args{123},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
