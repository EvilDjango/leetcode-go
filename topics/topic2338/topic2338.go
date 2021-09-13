// 面试题 17.22. 单词转换
//给定字典中的两个词，长度相等。写一个方法，把一个词转换成另一个词， 但是一次只能改变一个字符。每一步得到的新词都必须能在字典中找到。
//
//编写一个程序，返回一个可能的转换序列。如有多个可能的转换序列，你可以返回任何一个。
//
//示例 1:
//
//输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出:
//["hit","hot","dot","lot","log","cog"]
//示例 2:
//
//输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
//通过次数8,546提交次数22,946
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/13/21 6:02 PM
package topic2338

// 回溯
func findLadders(beginWord string, endWord string, wordList []string) []string {
	n := len(wordList)
	dict := make(map[string][]string, n)
	hasBeginWord := false
	for i := 0; i < n; i++ {
		w1 := wordList[i]
		if w1 == beginWord {
			hasBeginWord = true
		}

		for j := i + 1; j < n; j++ {
			w2 := wordList[j]
			if isAdjacent(w1, w2) {
				dict[w1] = append(dict[w1], w2)
				dict[w2] = append(dict[w2], w1)
			}
		}
	}
	if !hasBeginWord {
		for i := 0; i < n; i++ {
			w := wordList[i]
			if isAdjacent(beginWord, w) {
				dict[beginWord] = append(dict[beginWord], w)
				dict[w] = append(dict[w], beginWord)
			}
		}
	}
	return dfs(beginWord, endWord, dict, []string{beginWord}, make(map[string]bool, n))
}

func dfs(word string, endWord string, dict map[string][]string, words []string, used map[string]bool) []string {
	if word == endWord {
		return words
	}
	for _, w := range dict[word] {
		if used[w] {
			continue
		}
		words = append(words, w)
		used[w] = true
		res := dfs(w, endWord, dict, words, used)
		if res != nil {
			return res
		}
		words = words[:len(words)-1]
		// 如果运行到这一步，意味着无法从i这个点找到路径，used[w]无需改为false.这里解注释会超时。
		//used[w] = false
	}
	return nil
}
func isAdjacent(u, v string) bool {
	m, n := len(u), len(v)
	if m != n {
		return false
	}
	diff := 0
	for i := 0; i < m; i++ {
		if u[i] != v[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
