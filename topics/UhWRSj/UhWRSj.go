// 剑指 Offer II 063. 替换单词
//在英语中，有一个叫做 词根(root) 的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
//
//现在，给定一个由许多词根组成的词典和一个句子，需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
//
//需要输出替换之后的句子。
//
//
//
//示例 1：
//
//输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
//输出："the cat was rat by the bat"
//示例 2：
//
//输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
//输出："a a b c"
//示例 3：
//
//输入：dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
//输出："a a a a a a a a bbb baba a"
//示例 4：
//
//输入：dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
//输出："the cat was rat by the bat"
//示例 5：
//
//输入：dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
//输出："it is ab that this solution is ac"
//
//
//提示：
//
//1 <= dictionary.length <= 1000
//1 <= dictionary[i].length <= 100
//dictionary[i] 仅由小写字母组成。
//1 <= sentence.length <= 10^6
//sentence 仅由小写字母和空格组成。
//sentence 中单词的总量在范围 [1, 1000] 内。
//sentence 中每个单词的长度在范围 [1, 1000] 内。
//sentence 中单词之间由一个空格隔开。
//sentence 没有前导或尾随空格。
//
//
//注意：本题与主站 648 题相同： https://leetcode-cn.com/problems/replace-words/
//
//通过次数3,056提交次数4,228
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/29/21 9:20 PM
package UhWRSj

type trie struct {
	isEnd    bool
	children [26]*trie
}

func (t *trie) insert(s string) {
	curr := t
	for _, b := range s {
		i := b - 'a'
		if curr.children[i] == nil {
			curr.children[i] = &trie{}
		}
		curr = curr.children[i]
	}
	curr.isEnd = true
}

func (t *trie) find(s string) []byte {
	curr := t
	var bytes []byte
	for _, b := range s {
		curr = curr.children[b-'a']
		if curr == nil {
			break
		}
		bytes = append(bytes, byte(b))
		if curr.isEnd {
			return bytes
		}
	}
	return []byte(s)
}

func replaceWords(dictionary []string, sentence string) string {
	root := &trie{}
	for _, word := range dictionary {
		root.insert(word)
	}
	var bytes []byte
	for start, end := 0, 0; end <= len(sentence); end++ {
		if end == len(sentence) || sentence[end] == ' ' {
			bytes = append(bytes, root.find(sentence[start:end])...)
			bytes = append(bytes, ' ')
			end++
			start = end
		}
	}
	return string(bytes[:len(bytes)-1])
}
