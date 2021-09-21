// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/21/21 12:32 PM
package find_longest_subarray_lcci

import (
	"reflect"
	"testing"
)

func Test_findLongestSubarray(t *testing.T) {
	test_FindLongestSubarray(t, findLongestSubarray)
}

func Test_findLongestSubarray2(t *testing.T) {
	test_FindLongestSubarray(t, findLongestSubarray2)
}

func Test_findLongestSubarray3(t *testing.T) {
	test_FindLongestSubarray(t, findLongestSubarray3)
}

func test_FindLongestSubarray(t *testing.T, function func(array []string) []string) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}},
			[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7"},
		},
		{
			"",
			args{[]string{"A", "A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}},
			[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7"},
		},
		{
			"",
			args{[]string{"42", "10", "O", "t", "y", "p", "g", "B", "96", "H", "5", "v", "P", "52", "25", "96", "b", "L", "Y", "z", "d", "52", "3", "v", "71", "J", "A", "0", "v", "51", "E", "k", "H", "96", "21", "W", "59", "I", "V", "s", "59", "w", "X", "33", "29", "H", "32", "51", "f", "i", "58", "56", "66", "90", "F", "10", "93", "53", "85", "28", "78", "d", "67", "81", "T", "K", "S", "l", "L", "Z", "j", "5", "R", "b", "44", "R", "h", "B", "30", "63", "z", "75", "60", "m", "61", "a", "5", "S", "Z", "D", "2", "A", "W", "k", "84", "44", "96", "96", "y", "M"}},
			[]string{"52", "3", "v", "71", "J", "A", "0", "v", "51", "E", "k", "H", "96", "21", "W", "59", "I", "V", "s", "59", "w", "X", "33", "29", "H", "32", "51", "f", "i", "58", "56", "66", "90", "F", "10", "93", "53", "85", "28", "78", "d", "67", "81", "T", "K", "S", "l", "L", "Z", "j", "5", "R", "b", "44", "R", "h", "B", "30", "63", "z", "75", "60", "m", "61", "a", "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:  %v, \nwant: %v", got, tt.want)
			}
		})
	}
}
