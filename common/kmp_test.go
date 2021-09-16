// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/15/21 2:31 PM
package common

import (
	"strings"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		s   string
		sub string
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"",
			args{"abword", "word", 0, 6},
		},
		{
			"",
			args{"ABWORD", "WORD", 0, 6},
		},
		{
			"",
			args{"abword", "word1", 0, 6},
		},
		{
			"",
			args{"BBCABCDABABCDABCDABDE", "ABCDABC", 0, 21},
		},
		{
			"",
			args{"BBCABCDABABCDABCDABDE", "ABCDABD", 0, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strings.Index(tt.args.s, tt.args.sub)
			if got := search(tt.args.s, tt.args.sub, tt.args.l, tt.args.r); got != want {
				t.Errorf("search() = %v, want %v", got, want)
			}
		})
	}
}
