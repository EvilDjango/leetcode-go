// 面试题 17.25. 单词矩阵
//给定一份单词的清单，设计一个算法，创建由字母组成的面积最大的矩形，其中每一行组成一个单词(自左向右)，每一列也组成一个单词(自上而下)。不要求这些单词在清单里连续出现，但要求所有行等长，所有列等高。
//
//如果有多个面积最大的矩形，输出任意一个均可。一个单词可以重复使用。
//
//示例 1:
//
//输入: ["this", "real", "hard", "trh", "hea", "iar", "sld"]
//输出:
//[
//   "this",
//   "real",
//   "hard"
//]
//示例 2:
//
//输入: ["aa"]
//输出: ["aa","aa"]
//说明：
//
//words.length <= 1000
//words[i].length <= 100
//数据保证单词足够随机
//通过次数1,382提交次数2,803
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/9/21 2:46 PM
package topic2341_02

type TrieNode struct {
	IsLeaf   bool
	Children [26]*TrieNode
}

func (t *TrieNode) Insert(s string) {
	cur := t
	for _, r := range s {
		i := r - 'a'
		if cur.Children[i] == nil {
			cur.Children[i] = &TrieNode{}
		}
		cur = cur.Children[i]
	}
	cur.IsLeaf = true
}

// 利用前缀树来剪枝
func maxRectangle(words []string) []string {
	trie := &TrieNode{}
	wordsByLength := make(map[int][]string)
	maxLen := 0
	for _, word := range words {
		trie.Insert(word)
		l := len(word)
		wordsByLength[l] = append(wordsByLength[len(word)], word)
		if l > maxLen {
			maxLen = l
		}
	}

	tries := make([]*TrieNode, maxLen)
	for i := 0; i < len(tries); i++ {
		tries[i] = trie
	}
	var ans []string
	for l := maxLen; l > 0; l-- {
		if l*l <= area(ans) {
			break
		}
		if _, ok := wordsByLength[l]; !ok {
			continue
		}
		getRectangles(wordsByLength[l], 0, []string{}, tries, &ans)
	}
	return ans
}

func getRectangles(words []string, i int, rectangle []string, tries []*TrieNode, ans *[]string) {
	length := len(words[0])
	for _, word := range words {
		var nextTries []*TrieNode
		valid, allLeaf := true, true
		for i := 0; i < length; i++ {
			nextTrie := tries[i].Children[word[i]-'a']
			if nextTrie == nil {
				valid = false
				break
			}
			nextTries = append(nextTries, nextTrie)
			if !nextTrie.IsLeaf {
				allLeaf = false
			}
		}
		if valid {
			rectangle = append(rectangle, word)
			if allLeaf && area(rectangle) > area(*ans) {
				*ans = make([]string, len(rectangle))
				copy(*ans, rectangle)
			}
			if i+1 < length {
				getRectangles(words, i+1, rectangle, nextTries, ans)
			}
			rectangle = rectangle[:len(rectangle)-1]
		}
	}
}
func area(rectangle []string) int {
	if rectangle == nil {
		return 0
	}
	return len(rectangle) * len(rectangle[0])
}
