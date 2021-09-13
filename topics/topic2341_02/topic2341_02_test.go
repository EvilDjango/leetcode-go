// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/9/21 3:43 PM
package topic2341_02

import (
	"reflect"
	"testing"
)

func Test_maxRectangle(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{[]string{"this", "real", "hard", "trh", "hea", "iar", "sld"}},
			[]string{"trh", "hea", "iar", "sld"},
		},
		{
			"",
			args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[]string{"bat", "ate", "tea"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRectangle(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
