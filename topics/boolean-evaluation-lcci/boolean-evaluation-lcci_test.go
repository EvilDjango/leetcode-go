// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/15/21 2:14 PM
package boolean_evaluation_lcci

import "testing"

func Test_countEval(t *testing.T) {
	test(t, countEval)
}

func Test_countEval2(t *testing.T) {
	test(t, countEval2)
}

func test(t *testing.T, function func(s string, result int) int) {
	type args struct {
		s      string
		result int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"1^0|0|1", 0},
			2,
		},
		{
			"",
			args{"0&0&0&1^1|0", 1},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.s, tt.args.result); got != tt.want {
				t.Errorf("countEval() = %v, want %v", got, tt.want)
			}
		})
	}
}
