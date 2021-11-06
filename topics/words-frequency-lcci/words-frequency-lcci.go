// 面试题 16.02. 单词频率
//设计一个方法，找出任意指定单词在一本书中的出现频率。
//
//你的实现应该支持如下操作：
//
//WordsFrequency(book)构造函数，参数为字符串数组构成的一本书
//get(word)查询指定单词在书中出现的频率
//示例：
//
//WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
//wordsFrequency.get("you"); //返回0，"you"没有出现过
//wordsFrequency.get("have"); //返回2，"have"出现2次
//wordsFrequency.get("an"); //返回1
//wordsFrequency.get("apple"); //返回1
//wordsFrequency.get("pen"); //返回1
//提示：
//
//book[i]中只包含小写字母
//1 <= book.length <= 100000
//1 <= book[i].length <= 10
//get函数的调用次数不会超过100000
//通过次数15,252提交次数19,931
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/6/21 5:40 PM
package words_frequency_lcci

// 字典树解法
type Trie struct {
	count    int
	children [26]*Trie
}

func (t *Trie) insert(s string) {
	cur := t
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &Trie{}
		}
		cur = cur.children[index]
	}
	cur.count++
}

func (t *Trie) getCount(s string) int {
	cur := t
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if cur.children[index] == nil {
			return 0
		}
		cur = cur.children[index]
	}
	return cur.count
}

type WordsFrequency struct {
	trie *Trie
}

func Constructor(book []string) WordsFrequency {
	root := &Trie{}
	for _, b := range book {
		root.insert(b)
	}
	return WordsFrequency{root}
}

func (this *WordsFrequency) Get(word string) int {
	return this.trie.getCount(word)
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
