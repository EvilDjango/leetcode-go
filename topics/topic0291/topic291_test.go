package topic0291

import (
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/18/21 6:06 PM
*/

func Test_wordPatternMatch(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{"abab", "redblueredblue"},
			true,
		},
		{
			"",
			args{"aaaa", "asdasdasdasd"},
			true,
		},
		{
			"",
			args{"abab", "asdasdasdasd"},
			true,
		},
		{
			"",
			args{"aabb", "xyzabcxzyabc"},
			false,
		},
		{
			"",
			args{"abba", "dogcatcatdog"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPatternMatch(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPatternMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
