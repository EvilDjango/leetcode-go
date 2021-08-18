package topic269

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/3/21 11:28 PM
*/

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	"",
		//	args{[]string{"wrt", "wrf", "er", "ett", "rftt"}},
		//	"wertf",
		//},
		//{
		//	"",
		//	args{[]string{"z", "x"}},
		//	"zx",
		//},
		//{
		//	"",
		//	args{[]string{"z", "x","z"}},
		//	"",
		//},
		//{
		//	"",
		//	args{[]string{"z", "z"}},
		//	"z",
		//},
		//{
		//	"",
		//	args{[]string{"ab", "adc"}},
		//	"abcd",
		//},
		//{
		//	"",
		//	args{[]string{"ab", "ad","adc"}},
		//	"abcd",
		//},
		{
			"",
			args{[]string{"ab", "ad", "adb", "add", "addc"}},
			"abcd",
		},
		//{
		//	"",
		//	args{[]string{"abc", "ab"}},
		//	"",
		//},
		//{
		//	"",
		//	args{[]string{"blnw"}},
		//	"blnw",
		//},
		//
		//{
		//	"",
		//	args{[]string{"wnlb"}},
		//	"blnw",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
