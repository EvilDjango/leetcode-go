package topic291

/*
291. 单词规律 II
给你一种规律 pattern 和一个字符串 str，请你判断 str 是否遵循其相同的规律。

这里我们指的是 完全遵循，例如 pattern 里的每个字母和字符串 str 中每个 非空 单词之间，存在着 双射 的对应规律。双射 意味着映射双方一一对应，不会存在两个字符映射到同一个字符串，也不会存在一个字符分别映射到两个不同的字符串。



示例 1：

输入：pattern = "abab", s = "redblueredblue"
输出：true
解释：一种可能的映射如下：
'a' -> "red"
'b' -> "blue"
示例 2：

输入：pattern = "aaaa", s = "asdasdasdasd"
输出：true
解释：一种可能的映射如下：
'a' -> "asd"
示例 3：

输入：pattern = "abab", s = "asdasdasdasd"
输出：true
解释：一种可能的映射如下：
'a' -> "a"
'b' -> "sdasd"
注意 'a' 和 'b' 不能同时映射到 "asd"，因为这里的映射是一种双射。
示例 4：

输入：pattern = "aabb", s = "xyzabcxzyabc"
输出：false


提示：

0 <= pattern.length <= 20
0 <= s.length <= 50
pattern 和 s 由小写英文字母组成
通过次数2,359提交次数4,403

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/18/21 5:12 PM
*/
func wordPatternMatch(pattern string, s string) bool {
	p2s, s2p := make(map[byte]string), make(map[string]byte)
	return dfs(p2s, s2p, pattern, s, 0, 0)
}

func dfs(b2s map[byte]string, s2b map[string]byte, pattern, s string, i, j int) bool {
	pLen, sLen := len(pattern), len(s)
	if i == pLen || j == sLen {
		return i+j == pLen+sLen
	}
	b := pattern[i]
	for k := j + 1; k <= sLen; k++ {
		word := s[j:k]
		exist := b2s[b] == word
		if b2s[b] == "" && s2b[word] == 0 || b2s[b] == word {
			b2s[b] = word
			s2b[word] = b
			if dfs(b2s, s2b, pattern, s, i+1, k) {
				return true
			}
			if !exist {
				delete(b2s, b)
				delete(s2b, word)
			}
		}
	}
	return false
}
