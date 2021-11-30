// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/30/21 12:37 PM
package bianry_number_to_string_lcci

import "testing"

func Test_printBin(t *testing.T) {
	test(t, printBin)
}

func Test_printBin2(t *testing.T) {
	test(t, printBin2)
}

func test(t *testing.T, function func(num float64) string) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{0.625},
			"0.101",
		},
		{
			"",
			args{0.1},
			"ERROR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.num); got != tt.want {
				t.Errorf("printBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
