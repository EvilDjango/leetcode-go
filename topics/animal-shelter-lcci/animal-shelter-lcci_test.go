// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/3/21 4:04 PM
package animal_shelter_lcci

import (
	"leetcode-go/common"
	"testing"
)

func TestAnimalShelf_DequeueAny(t *testing.T) {
	nothing := []int{-1, -1}
	shelf := Constructor()
	shelf.Enqueue([]int{0, 0})
	shelf.Enqueue([]int{1, 0})
	t.Run("", func(t *testing.T) {
		res := shelf.DequeueCat()
		want := []int{0, 0}
		if !common.ArrayEqual(res, want) {
			t.Errorf("DequeueDog() returned %v, want %v", res, want)
		}
	})
	t.Run("", func(t *testing.T) {
		res := shelf.DequeueDog()
		if !common.ArrayEqual(res, nothing) {
			t.Errorf("DequeueDog() returned %v, want %v", res, nothing)
		}
	})
	t.Run("", func(t *testing.T) {
		res := shelf.DequeueAny()
		want := []int{1, 0}
		if !common.ArrayEqual(res, want) {
			t.Errorf("DequeueDog() returned %v, want %v", res, want)
		}
	})

}
