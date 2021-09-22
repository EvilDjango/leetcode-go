// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/22/21 9:13 PM
package langtons_ant_lcci_02

import (
	"reflect"
	"testing"
)

func Test_printKMoves(t *testing.T) {
	test(t, printKMoves)
}

func Test_printKMoves2(t *testing.T) {
	test(t, printKMoves2)
}

func test(t *testing.T, function func(K int) []string) {
	type args struct {
		K int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{0},
			[]string{"R"},
		},
		{
			"",
			args{2},
			[]string{"_X", "LX"},
		},
		{
			"",
			args{5},
			[]string{"_U", "X_", "XX"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printKMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
