// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/10/21 2:29 PM
package re_space_lcci

import "testing"

func TestStreamRank_GetRankOfNumber(t *testing.T) {
	rank := Constructor()
	rank.Track(0)
	t.Run("", func(t *testing.T) {
		if got := rank.GetRankOfNumber(0); got != 1 {
			t.Errorf("GetRankOfNumber() = %v, want %v", got, 1)
		}
	})
	rank.Track(1)
	t.Run("", func(t *testing.T) {
		if got := rank.GetRankOfNumber(1); got != 2 {
			t.Errorf("GetRankOfNumber() = %v, want %v", got, 2)
		}
	})
}
