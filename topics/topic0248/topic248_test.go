package topic0248

import (
	"reflect"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/22/21 5:24 PM
*/

func Test_compare(t *testing.T) {
	type args struct {
		i string
		j string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{"0", "0"},
			0,
		},
		{
			"case2",
			args{"1", "0"},
			1,
		},
		{
			"case3",
			args{"123", "122"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strobogrammatic(t *testing.T) {
	type args struct {
		len    int
		remain int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{1, 1},
			[]string{"0", "1", "8"},
		},
		{
			"",
			args{0, 0},
			[]string{""},
		},
		{
			"",
			args{2, 2},
			[]string{"11", "69", "88", "96"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strobogrammatic(tt.args.len, tt.args.remain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strobogrammatic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strobogrammaticInRange(t *testing.T) {
	type args struct {
		low  string
		high string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"",
			args{"0", "0"},
			1,
		},
		{"",
			args{"0", "1"},
			2,
		},
		{"",
			args{"0", "100"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strobogrammaticInRange(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("strobogrammaticInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
