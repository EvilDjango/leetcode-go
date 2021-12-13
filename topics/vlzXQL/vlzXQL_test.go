// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/13/21 3:53 PM
package vlzXQL

import (
	"leetcode-go/common"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	test(t, calcEquation)
}

func Test_calcEquation2(t *testing.T) {
	test(t, calcEquation2)
}

func test(t *testing.T, function func(equations [][]string, values []float64, queries [][]string) []float64) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{

		{
			"",
			args{
				[][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}},
				[]float64{3.4, 1.4, 2.3},
				[][]string{{"b", "a"}, {"a", "f"}, {"f", "f"}, {"e", "e"}, {"c", "c"}, {"a", "c"}, {"f", "e"}},
			},
			[]float64{0.29412, 10.948, 1.0, 1.0, -1.0, -1.0, 0.71429},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := function(tt.args.equations, tt.args.values, tt.args.queries)
			for i := 0; i < len(got); i++ {
				if !common.FloatEquals(got[i], tt.want[i], 5) {
					t.Errorf("calcEquation(%v) = %v, want %v", tt.args.queries[i], got[i], tt.want[i])
				}
			}
		})
	}
}
