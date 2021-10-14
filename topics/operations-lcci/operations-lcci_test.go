// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/14/21 9:17 AM
package operations_lcci

import "testing"

func TestOperations_opposite(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{0},
			0,
		},
		{
			"",
			args{1},
			-1,
		},
		{
			"",
			args{-1},
			1,
		},
		{
			"",
			args{1<<31 - 1},
			-(1<<31 - 1),
		},
		{
			"",
			args{-1 << 31},
			1 << 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			if got := this.opposite(tt.args.a); got != tt.want {
				t.Errorf("opposite() = %v, want %v", got, tt.want)
			}
		})
	}
}
