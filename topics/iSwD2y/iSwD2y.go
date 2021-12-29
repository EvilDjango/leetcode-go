// 剑指 Offer II 065. 最短的单词编码
//单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：
//
//words.length == indices.length
//助记字符串 s 以 '#' 字符结尾
//对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
//给定一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。
//
//
//
//示例 1：
//
//输入：words = ["time", "me", "bell"]
//输出：10
//解释：一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。
//words[0] = "time" ，s 开始于 indices[0] = 0 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
//words[1] = "me" ，s 开始于 indices[1] = 2 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
//words[2] = "bell" ，s 开始于 indices[2] = 5 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
//示例 2：
//
//输入：words = ["t"]
//输出：2
//解释：一组有效编码为 s = "t#" 和 indices = [0] 。
//
//
//提示：
//
//1 <= words.length <= 2000
//1 <= words[i].length <= 7
//words[i] 仅由小写字母组成
//
//
//注意：本题与主站 820 题相同： https://leetcode-cn.com/problems/short-encoding-of-words/
//
//通过次数2,302提交次数3,555
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/29/21 2:46 PM
package iSwD2y

import (
	"sort"
	"strings"
)

type trie struct {
	hasChild bool
	children [26]*trie
}

func (t *trie) insert(s string) *trie {
	curr := t
	for i := len(s) - 1; i >= 0; i-- {
		j := s[i] - 'a'
		if curr.children[j] == nil {
			curr.hasChild = true
			curr.children[j] = &trie{}
		}
		curr = curr.children[j]
	}
	return curr
}

func minimumLengthEncoding(words []string) int {
	root := &trie{}
	ends := map[*trie]int{}
	for i, word := range words {
		ends[root.insert(word)] = i
	}
	ans := 0
	for t, i := range ends {
		if !t.hasChild {
			ans += len(words[i]) + 1
		}
	}
	return ans
}

// 参考官方题解，存储末位节点，省去了深度搜索
func minimumLengthEncoding2(words []string) int {
	root := &trie{}
	ends := map[*trie]int{}
	for i, word := range words {
		ends[root.insert(word)] = i
	}
	ans := 0
	for t, i := range ends {
		if !t.hasChild {
			ans += len(words[i]) + 1
		}
	}
	return ans
}

// 按后缀排序
func minimumLengthEncoding3(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return compare(words[i], words[j])
	})
	ans := 0
	for i := 0; i < len(words); i++ {
		if i+1 == len(words) || !strings.HasSuffix(words[i+1], words[i]) {
			ans += len(words[i]) + 1
		}
	}
	return ans
}

func compare(a, b string) bool {
	l1, l2 := len(a), len(b)
	l := l1
	if l2 < l {
		l = l2
	}
	for i := 0; i < l; i++ {
		if a[l1-1-i] != b[l2-1-i] {
			return a[l1-1-i] < b[l2-1-i]
		}
	}
	return l1 < l2
}
