// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/14/21 10:19 PM
package zlDJc7

import "testing"

func Test_openLock(t *testing.T) {
	test(t, openLock)
}

func Test_openLock2(t *testing.T) {
	test(t, openLock2)
}

func Test_openLock3(t *testing.T) {
	test(t, openLock3)
}

func test(t *testing.T, function func(deadends []string, target string) int) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]string{}, "1000"},
			1,
		},
		{
			"",
			args{[]string{}, "1900"},
			2,
		},
		{
			"",
			args{[]string{}, "1111"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
