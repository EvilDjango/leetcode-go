package topics

import "math"

/*
301. 删除无效的括号
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回。



示例 1：

输入：s = "()())()"
输出：["(())()","()()()"]
示例 2：

输入：s = "(a)())()"
输出：["(a())()","(a)()()"]
示例 3：

输入：s = ")("
输出：[""]


提示：

1 <= s.length <= 25
s 由小写英文字母以及括号 '(' 和 ')' 组成
s 中至多含 20 个括号
通过次数30,532提交次数58,480

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/24/21 1:12 PM
*/

func removeInvalidParentheses(s string) []string {
	var ans []string
	minDelete := math.MaxInt32
	backtrack(s, &ans, 0, 0, 0, "", &minDelete, false)
	return ans
}

func backtrack(s string, ans *[]string, i, sum, deleted int, sub string, minDelete *int, prevDelete bool) {
	if sum == -1 || deleted > *minDelete {
		return
	}
	if i == len(s) {
		if sum != 0 {
			return
		}
		if deleted == *minDelete {
			*ans = append(*ans, sub)
		} else {
			*ans = []string{sub}
			*minDelete = deleted
		}
		return
	}

	// 移除当前位置的字符，若当前字符和前面的字符相同，则仅当前面的字符被移除时才考虑移除当前字符，这样可以避免重复
	if (s[i] == '(' || s[i] == ')') && (prevDelete || i == 0 || s[i] != s[i-1]) {
		backtrack(s, ans, i+1, sum, deleted+1, sub, minDelete, true)
	}

	// 保留当前字符
	if s[i] == '(' {
		sum++
	} else if s[i] == ')' {
		sum--
	}
	backtrack(s, ans, i+1, sum, deleted, sub+s[i:i+1], minDelete, false)
}
