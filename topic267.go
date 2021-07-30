package leetcode_go

/*
267. 回文排列 II
给定一个字符串 s ，返回其通过重新排列组合后所有可能的回文字符串，并去除重复的组合。

如不能形成任何回文排列时，则返回一个空列表。

示例 1：

输入: "aabb"
输出: ["abba", "baab"]
示例 2：

输入: "abc"
输出: []
通过次数3,030提交次数6,907

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/29/21 11:59 AM
*/

func generatePalindromes(s string) []string {
	cnt := make(map[rune]int, len(s))
	for _, r := range s {
		cnt[r]++
	}
	center := ""
	for r, n := range cnt {
		if n%2 == 0 {
			continue
		}
		if center == "" {
			center = string(r)
		} else {
			// 有两个以上的奇数字符，无法组成回文串
			return nil
		}
	}
	if center != "" {
		cnt[rune(center[0])]--
	}
	var ans []string
	dfs(cnt, len(s), center, &ans)
	return ans
}

func dfs(dic map[rune]int, n int, s string, strs *[]string) {
	if len(s) == n {
		*strs = append(*strs, s)
		return
	}
	for r, c := range dic {
		if c > 0 {
			dic[r] -= 2
			dfs(dic, n, string(r)+s+string(r), strs)
			dic[r] += 2
		}
	}
}
