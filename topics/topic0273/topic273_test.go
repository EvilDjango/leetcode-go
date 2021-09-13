package topic0273

import "testing"

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/6/21 3:28 PM
*/

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{0},
			"Zero",
		},
		{
			"",
			args{10},
			"Ten",
		},
		{
			"",
			args{30},
			"Thirty",
		},
		{
			"",
			args{100},
			"One Hundred",
		},
		{
			"",
			args{103},
			"One Hundred Three",
		},
		{
			"",
			args{110},
			"One Hundred Ten",
		},
		{
			"",
			args{111},
			"One Hundred Eleven",
		},
		{
			"",
			args{123},
			"One Hundred Twenty Three",
		},
		{
			"",
			args{12345},
			"Twelve Thousand Three Hundred Forty Five",
		},
		{
			"",
			args{1234567},
			"One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			"",
			args{1234567891},
			"One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
