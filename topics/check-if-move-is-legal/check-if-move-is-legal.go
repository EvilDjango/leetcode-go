// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/7/7 上午11:11
package check_if_move_is_legal

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	opposite := byte('W')
	if color == 'W' {
		opposite = byte('B')
	}
	f := func(r, c int) bool {
		nextR := rMove
		nextC := cMove
		found := false
		for {
			nextR += r
			nextC += c
			if nextR >= len(board) || nextC >= len(board[0]) {
				return false
			}
			curr := board[nextR][nextC]
			if !found {
				if curr == opposite {
					found = true
				} else {
					return false
				}
			} else {
				if curr == '.' {
					return false
				}
				if curr == color {
					return true
				}
			}
		}
	}
	directions := [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	for _, direction := range directions {
		if f(direction[0], direction[1]) {
			return true
		}
	}
	return false
}
