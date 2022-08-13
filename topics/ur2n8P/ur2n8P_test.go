// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/7/23 下午10:42
package ur2n8P

import "testing"

func Test_sequenceReconstruction2(t *testing.T) {
	test(t, sequenceReconstruction2)
}

func test(t *testing.T, function func(org []int, seqs [][]int) bool) {
	type args struct {
		org  []int
		seqs [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{
				[]int{1, 2, 3},
				[][]int{
					{1, 2},
					{1, 3},
					{2, 3},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.org, tt.args.seqs); got != tt.want {
				t.Errorf("sequenceReconstruction2() = %v, want %v", got, tt.want)
			}
		})
	}
}
