package topic0294

/*
294. 翻转游戏 II
你和朋友玩一个叫做「翻转游戏」的游戏。游戏规则如下：

给你一个字符串 currentState ，其中只含 '+' 和 '-' 。你和朋友轮流将 连续 的两个 "++" 反转成 "--" 。当一方无法进行有效的翻转时便意味着游戏结束，则另一方获胜。

请你写出一个函数来判定起始玩家 是否存在获胜的方案 ：如果存在，返回 true ；否则，返回 false 。


示例 1：

输入：currentState = "++++"
输出：true
解释：起始玩家可将中间的 "++" 翻转变为 "+--+" 从而得胜。
示例 2：

输入：currentState = "+"
输出：false


提示：

1 <= currentState.length <= 60
currentState[i] 不是 '+' 就是 '-'


进阶：请推导你算法的时间复杂度。

通过次数3,365提交次数5,745

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/19/21 11:31 AM
*/

func canWin(currentState string) bool {
	cache := make(map[string]bool)
	return dfs(cache, currentState)
}

func dfs(cache map[string]bool, state string) bool {
	if win, ok := cache[state]; ok {
		return win
	}
	for j := 1; j < len(state); j++ {
		if state[j-1:j+1] == "++" {
			newState := state[:j-1] + "--" + state[j+1:]
			if !dfs(cache, newState) {
				cache[state] = true
				return true
			}
		}
	}
	cache[state] = false
	return false
}
