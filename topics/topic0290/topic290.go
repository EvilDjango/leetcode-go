package topic0290

import "strings"

/*
290. 单词规律
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

通过次数78,456提交次数171,648

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/18/21 3:58 PM
*/

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	b2s := make(map[byte]string)
	s2b := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		b, word := pattern[i], words[i]
		if v, ok := b2s[b]; ok && v != word {
			return false
		}
		if v, ok := s2b[word]; ok && v != b {
			return false
		}
		b2s[b] = word
		s2b[word] = b
	}
	return true
}
