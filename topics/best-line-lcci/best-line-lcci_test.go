// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/25/21 10:39 PM
package best_line_lcci

import (
	"reflect"
	"testing"
)

func Test_bestLine(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[][]int{{26072, -12996}, {-41195, -34139}, {6491, 14145}, {275, 4007}, {14321, -15055}, {-38983, -49757}, {-28710, -15391}, {-29300, 12859}, {-34606, -25274}, {-37657, 14795}, {-32300, 1599}, {-24623, -14921}, {-35555, -43348}, {-41350, -16826}, {-38848, -6454}, {5588, -6650}, {-8414, -38364}, {15409, 20374}, {29264, 28598}, {-9066, -388}, {-32891, -25982}, {4402, 6766}, {-12017, -22656}, {-35555, -12886}, {-10073, -11487}, {10118, -18119}, {-11704, 11154}, {-25250, 28060}, {-36845, -27142}, {-16539, -8717}, {-9274, 23635}, {-7038, -17688}, {-4654, -3477}, {-30050, 10044}, {-31933, -42528}, {-20460, -15066}, {27274, -18550}, {22048, 28678}, {-35555, -17101}, {-33957, 26896}, {-8262, -11077}, {15830, -9823}, {-38355, -30257}, {14949, 4445}, {-24900, 21759}, {-24800, 29749}, {-35555, -18966}, {10459, 17639}, {-40180, -29454}, {-24194, -17257}, {-9540, -28060}, {-8029, -4150}}},
			[]int{2, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestLine(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestLine() = %v, want %v", got, tt.want)
			}
		})
	}
}