// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/14/21 10:47 AM
package english_int_lcci

import "testing"

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
			args{12345},
			"Twelve Thousand Three Hundred Forty Five",
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
