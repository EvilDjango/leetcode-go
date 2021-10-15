// 面试题 16.04. 井字游戏
//设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。
//
//以下是井字游戏的规则：
//
//玩家轮流将字符放入空位（" "）中。
//第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
//"X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
//当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
//当所有位置非空时，也算为游戏结束。
//如果游戏结束，玩家不允许再放置字符。
//如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。
//
//示例 1：
//
//输入： board = ["O X"," XO","X O"]
//输出： "X"
//示例 2：
//
//输入： board = ["OOX","XXO","OXO"]
//输出： "Draw"
//解释： 没有玩家获胜且不存在空位
//示例 3：
//
//输入： board = ["OOX","XXO","OX "]
//输出： "Pending"
//解释： 没有玩家获胜且仍存在空位
//提示：
//
//1 <= board.length == board[i].length <= 100
//输入一定遵循井字棋规则
//通过次数7,217提交次数15,697
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/15/21 9:37 AM
package tic_tac_toe_lcci

func tictactoe(board []string) string {
	N := len(board)
	winner, finish := uint8(' '), true
	// 检查行
	for i := 0; i < N; i++ {
		winner = board[i][0]
		for j := 0; j < N; j++ {
			if j > 0 && board[i][j] != board[i][j-1] {
				winner = ' '
			}
			if board[i][j] == ' ' {
				finish = false
				break
			}
		}
		if winner != ' ' {
			return string(winner)
		}
	}

	// 检查列
	for i := 0; i < N; i++ {
		winner = board[0][i]
		for j := 1; j < N; j++ {
			if board[j][i] != board[j-1][i] {
				winner = ' '
			}
		}
		if winner != ' ' {
			return string(winner)
		}
	}

	//检查左上到右下的对角线
	winner = board[0][0]
	for i := 0; i < N-1; i++ {
		if board[i][i] != board[i+1][i+1] {
			winner = ' '
			break
		}
	}
	if winner != ' ' {
		return string(winner)
	}

	//检查左下到右上的对角线
	winner = board[N-1][0]
	for i := N - 1; i > 0; i-- {
		if board[i][N-1-i] != board[i-1][N-i] {
			winner = ' '
			break
		}
	}
	if winner != ' ' {
		return string(winner)
	}

	if finish {
		return "Draw"
	}
	return "Pending"
}
