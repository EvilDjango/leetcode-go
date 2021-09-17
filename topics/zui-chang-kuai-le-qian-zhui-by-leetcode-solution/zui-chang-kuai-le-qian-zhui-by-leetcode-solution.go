// 1392. 最长快乐前缀
//「快乐前缀」是在原字符串中既是 非空 前缀也是后缀（不包括原字符串自身）的字符串。
//
//给你一个字符串 s，请你返回它的 最长快乐前缀。
//
//如果不存在满足题意的前缀，则返回一个空字符串。
//
//
//
//示例 1：
//
//输入：s = "level"
//输出："l"
//解释：不包括 s 自己，一共有 4 个前缀（"l", "le", "lev", "leve"）和 4 个后缀（"l", "el", "vel", "evel"）。最长的既是前缀也是后缀的字符串是 "l" 。
//示例 2：
//
//输入：s = "ababab"
//输出："abab"
//解释："abab" 是最长的既是前缀也是后缀的字符串。题目允许前后缀在原字符串中重叠。
//示例 3：
//
//输入：s = "leetcodeleet"
//输出："leet"
//示例 4：
//
//输入：s = "a"
//输出：""
//
//
//提示：
//
//1 <= s.length <= 10^5
//s 只含有小写英文字母
//通过次数7,529提交次数18,769
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 11:37 AM
package zui_chang_kuai_le_qian_zhui_by_leetcode_solution

// kmp算法
func longestPrefix(s string) string {
	size := len(s)
	next := make([]int, size+1)
	next[0] = -1
	i, j := 0, -1
	for i < size && j < size {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return s[size-next[size]:]
}

// kmp算法的另一种写法
func longestPrefix2(s string) string {
	size := len(s)
	p := make([]int, size)
	i, j := 1, 0
	for i < size && j < size {
		if s[i] == s[j] {
			j++
			p[i] = j
			i++
		} else if j == 0 {
			i++
		} else {
			j = p[j-1]
		}
	}
	return s[size-p[size-1]:]
}

// kmp算法的另一种写法
func longestPrefix3(s string) string {
	size := len(s)
	fail := make([]int, size)
	for i := 1; i < size; i++ {
		j := fail[i-1]
		for j != 0 && s[i] != s[j] {
			j = fail[j-1]
		}
		if s[i] == s[j] {
			fail[i] = j + 1
		}
	}
	return s[size-fail[size-1]:]
}

// Rabin-Karp算法
func longestPrefix4(s string) string {
	base := int64(31)
	highBase := int64(1)
	length := 0
	prefix, suffix := int64(0), int64(0)
	for l := 1; l < len(s); l++ {
		prefix = prefix*base + int64(s[l-1]-'a')
		suffix = highBase*int64(s[len(s)-l]-'a') + suffix
		highBase *= base
		if prefix == suffix {
			length = l
		}
	}
	return s[:length]
}
