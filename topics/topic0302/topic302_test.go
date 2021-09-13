package topic0302

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/24/21 5:52 PM
*/

func Test_minArea(t *testing.T) {
	type args struct {
		image [][]byte
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[][]byte{{'0', '0', '1', '0'}, {'0', '1', '1', '0'}, {'0', '1', '0', '0'}}, 0, 2},
			6,
		},
		{
			"",
			args{[][]byte{{'1'}}, 0, 0},
			1,
		},
		{
			"",
			args{[][]byte{{'0', '1', '0'}, {'0', '1', '1'}, {'0', '0', '0'}}, 0, 1},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArea(tt.args.image, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("minArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
