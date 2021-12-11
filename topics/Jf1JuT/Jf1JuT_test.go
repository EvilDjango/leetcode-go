// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/11/21 1:09 PM
package Jf1JuT

import "testing"

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{[]string{"wrt", "wrf", "er", "ett", "rftt"}},
			"wertf",
		},
		{
			"",
			args{[]string{"zy", "zx"}},
			"yzx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
