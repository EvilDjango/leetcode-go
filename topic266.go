package leetcode_go

/*
266. 回文排列
给定一个字符串，判断该字符串中是否可以通过重新排列组合，形成一个回文字符串。

示例 1：

输入: "code"
输出: false
示例 2：

输入: "aab"
输出: true
示例 3：

输入: "carerac"
输出: true
通过次数5,947提交次数8,884

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 11:22 PM
*/

func canPermutePalindrome(s string) bool {
	set := make(map[rune]bool)
	for _, r := range s {
		if set[r] {
			delete(set, r)
		} else {
			set[r] = true
		}
	}
	return len(set) < 2
}
