// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 11:13 PM
package string_to_url_lcci

import "testing"

func Test_replaceSpaces(t *testing.T) {
	test(t, replaceSpaces)
}

func Test_replaceSpaces2(t *testing.T) {
	test(t, replaceSpaces2)
}

func test(t *testing.T, function func(S string, length int) string) {
	type args struct {
		S      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{"ds sdfs afs sdfa dfssf asdf             ", 27},
			"ds%20sdfs%20afs%20sdfa%20dfssf%20asdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := function(tt.args.S, tt.args.length); got != tt.want {
				t.Errorf("replaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
