// 面试题 17.15. 最长单词
//给定一组单词words，编写一个程序，找出其中的最长单词，且该单词由这组单词中的其他单词组合而成。若有多个长度相同的结果，返回其中字典序最小的一项，若没有符合要求的单词则返回空字符串。
//
//示例：
//
//输入： ["cat","banana","dog","nana","walk","walker","dogwalker"]
//输出： "dogwalker"
//解释： "dogwalker"可由"dog"和"walker"组成。
//提示：
//
//0 <= len(words) <= 200
//1 <= len(words[i]) <= 100
//通过次数5,668提交次数14,077
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/16/21 4:12 PM
package topic2342

type Trie struct {
	IsLeaf   bool
	Children [26]*Trie
}

func (t *Trie) Insert(s string) {
	curr := t
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if curr.Children[index] == nil {
			curr.Children[index] = &Trie{}
		}
		curr = curr.Children[index]
	}
	curr.IsLeaf = true
}

func (t *Trie) Find(b byte) *Trie {
	return t.Children[b-'a']
}

func longestWord(words []string) string {
	root := &Trie{}
	for _, word := range words {
		root.Insert(word)
	}
	ans := ""
	for _, word := range words {
		if (len(word) > len(ans) || len(word) == len(ans) && word < ans) && isCombined(root, root, word, 0, 0) {
			ans = word
		}
	}
	return ans
}

func isCombined(root, curr *Trie, word string, i, cnt int) bool {
	if i == len(word) {
		if curr == root && cnt > 1 {
			return true
		}
		return false
	}
	curr = curr.Find(word[i])
	if curr == nil {
		return false
	}
	if curr.IsLeaf && isCombined(root, root, word, i+1, cnt+1) {
		return true
	}
	return isCombined(root, curr, word, i+1, cnt)
}
