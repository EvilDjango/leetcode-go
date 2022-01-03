// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/31/21 2:51 PM
package fi9suh

import "testing"

func TestMyCalendar_Book(t *testing.T) {
	calender := Constructor()
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{47, 50},
			true,
		},
		{
			"",
			args{33, 41},
			true,
		},
		{
			"",
			args{39, 45},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calender.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
