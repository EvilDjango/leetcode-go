// 剑指 Offer II 062. 实现前缀树
//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
//
//请你实现 Trie 类：
//
//Trie() 初始化前缀树对象。
//void insert(String word) 向前缀树中插入字符串 word 。
//boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
//boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
//
//
//示例：
//
//输入
//inputs = ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//inputs = [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
//
//解释
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 True
//trie.search("app");     // 返回 False
//trie.startsWith("app"); // 返回 True
//trie.insert("app");
//trie.search("app");     // 返回 True
//
//
//提示：
//
//1 <= word.length, prefix.length <= 2000
//word 和 prefix 仅由小写英文字母组成
//insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
//
//
//
//
//注意：本题与主站 208 题相同：https://leetcode-cn.com/problems/implement-trie-prefix-tree/
//
//通过次数3,756提交次数4,879
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/30/21 12:16 PM
package QC3q1f

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, b := range word {
		i := b - 'a'
		if curr.children[i] == nil {
			curr.children[i] = &Trie{}
		}
		curr = curr.children[i]
	}
	curr.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	curr := this
	for _, b := range prefix {
		i := b - 'a'
		if curr.children[i] == nil {
			return nil
		}
		curr = curr.children[i]
	}
	return curr
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
