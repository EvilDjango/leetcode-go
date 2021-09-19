// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/20/21 2:07 AM
package baby_names_lcci_02

import (
	"reflect"
	"testing"
)

func Test_trulyMostPopular(t *testing.T) {
	type args struct {
		names    []string
		synonyms []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{
				[]string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"},
				[]string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"},
			},
			[]string{"John(27)", "Chris(36)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trulyMostPopular(tt.args.names, tt.args.synonyms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trulyMostPopular() = %v, want %v", got, tt.want)
			}
		})
	}
}
