package common

import (
	"reflect"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 12:56 PM
*/

func TestLowerBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 2, 3, 4}, 3},
			2,
		},
		{
			"",
			args{[]int{1, 2, 3, 4}, 4},
			3,
		},
		{
			"",
			args{[]int{1, 2, 3, 4}, 5},
			4,
		},
		{
			"",
			args{[]int{1, 2, 3, 4}, 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.nums, float64(tt.args.target)); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerBoundPart(t *testing.T) {
	type args struct {
		nums   []int
		target int
		l      int
		r      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 2, 3, 4, 5, 6}, 4, 0, 3},
			3,
		},
		{
			"",
			args{[]int{1, 2, 3, 4, 4, 6}, 3, 0, 5},
			2,
		},
		{
			"",
			args{[]int{1, 2, 3, 4, 5, 6}, 4, 1, 3},
			3,
		},
		{
			"",
			args{[]int{1, 2, 3, 4, 5, 6}, 4, 4, 6},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBoundPart(tt.args.nums, float64(tt.args.target), tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("LowerBoundPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 2, 2, 3, 4}, 2},
			3,
		},
		{
			"",
			args{[]int{1, 2, 2, 3, 4}, 0},
			0,
		},
		{
			"",
			args{[]int{1, 2, 2, 3, 4}, 5},
			5,
		},
		{
			"",
			args{[]int{1, 2, 2, 3, 4}, 3},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperBound(tt.args.nums, float64(tt.args.target)); got != tt.want {
				t.Errorf("UpperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBoundPart(t *testing.T) {
	type args struct {
		nums   []int
		target int
		l      int
		r      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 2, 2, 3, 3, 4, 5}, 3, 0, 2},
			2,
		},
		{
			"",
			args{[]int{1, 2, 2, 3, 3, 4, 5}, 3, 3, 5},
			5,
		},
		{
			"",
			args{[]int{1, 2, 2, 3, 3, 4, 5}, 2, 2, 4},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperBoundPart(tt.args.nums, float64(tt.args.target), tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("UpperBoundPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxWithoutOne(t *testing.T) {
	testMaxWithoutOne(t, MaxWithoutOne)
}

func TestMaxWithoutOne2(t *testing.T) {
	testMaxWithoutOne(t, MaxWithoutOne2)
}
func testMaxWithoutOne(t *testing.T, f func([]int) []int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{5, 5, 5, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxWithoutOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinWithoutOne(t *testing.T) {
	testMinWithoutOne(t, MinWithoutOne)
}

func TestMinWithoutOne2(t *testing.T) {
	testMinWithoutOne(t, MinWithoutOne2)
}

func testMinWithoutOne(t *testing.T, f func([]int) []int) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[]int{5, 4, 3, 2, 1}},
			[]int{1, 1, 1, 1, 2},
		},
		{
			"",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{2, 1, 1, 1, 1},
		},
		{
			"",
			args{[]int{1}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinWithoutOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
