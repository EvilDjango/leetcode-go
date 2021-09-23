// 面试题 16.20. T9键盘
//在老式手机上，用户通过数字键盘输入，手机将提供与这些数字相匹配的单词列表。每个数字映射到0至4个字母。给定一个数字序列，实现一个算法来返回匹配单词的列表。你会得到一张含有有效单词的列表。映射如下图所示：
//
//
//
//示例 1:
//
//输入: num = "8733", words = ["tree", "used"]
//输出: ["tree", "used"]
//示例 2:
//
//输入: num = "2", words = ["a", "b", "c", "d"]
//输出: ["a", "b", "c"]
//提示：
//
//num.length <= 1000
//words.length <= 500
//words[i].length == num.length
//num中不会出现 0, 1 这两个数字
//通过次数8,665提交次数12,020
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/23/21 1:21 PM
package t9_lcci

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
func (t *Trie) Search(b byte) *Trie {
	return t.Children[b-'a']
}

// 字典树的普通应用罢了
func getValidT9Words(num string, words []string) []string {
	dict := map[byte][]byte{}
	dict['2'] = []byte{'a', 'b', 'c'}
	dict['3'] = []byte{'d', 'e', 'f'}
	dict['4'] = []byte{'g', 'h', 'i'}
	dict['5'] = []byte{'j', 'k', 'l'}
	dict['6'] = []byte{'m', 'n', 'o'}
	dict['7'] = []byte{'p', 'q', 'r', 's'}
	dict['8'] = []byte{'t', 'u', 'v'}
	dict['9'] = []byte{'w', 'x', 'y', 'z'}
	root := &Trie{}
	for _, word := range words {
		root.Insert(word)
	}
	var ans []string
	var dfs func(index int, t *Trie, word []byte)
	dfs = func(index int, t *Trie, word []byte) {
		if index == len(num) {
			if t.IsLeaf {
				ans = append(ans, string(word))
			}
			return
		}
		for _, b := range dict[num[index]] {
			next := t.Search(b)
			if next == nil {
				continue
			}
			word = append(word, b)
			dfs(index+1, next, word)
			word = word[:len(word)-1]
		}
	}
	dfs(0, root, []byte{})
	return ans
}

// 暴力回溯，使用了哈希表,会超时
func getValidT9Words2(num string, words []string) []string {
	keyboard := map[byte][]byte{}
	keyboard['2'] = []byte{'a', 'b', 'c'}
	keyboard['3'] = []byte{'d', 'e', 'f'}
	keyboard['4'] = []byte{'g', 'h', 'i'}
	keyboard['5'] = []byte{'j', 'k', 'l'}
	keyboard['6'] = []byte{'m', 'n', 'o'}
	keyboard['7'] = []byte{'p', 'q', 'r', 's'}
	keyboard['8'] = []byte{'t', 'u', 'v'}
	keyboard['9'] = []byte{'w', 'x', 'y', 'z'}
	dict := make(map[string]bool, len(words))
	for _, word := range words {
		dict[word] = true
	}
	var ans []string
	var dfs func(index int, bytes []byte)
	dfs = func(index int, bytes []byte) {
		if index == len(num) {
			word := string(bytes)
			if dict[word] {
				ans = append(ans, word)
			}
			return
		}
		for _, b := range keyboard[num[index]] {
			bytes = append(bytes, b)
			dfs(index+1, bytes)
			bytes = bytes[:len(bytes)-1]
		}
	}
	dfs(0, []byte{})
	return ans
}

// 参考了优秀解法
func getValidT9Words3(num string, words []string) []string {
	var ans []string
	for _, word := range words {
		i := 0
		for ; i < len(word); i++ {
			if "22233344455566677778889999"[word[i]-'a'] != num[i] {
				break
			}
		}
		if i == len(word) {
			ans = append(ans, word)
		}
	}
	return ans
}
