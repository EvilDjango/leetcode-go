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

// 利用bellman-ford算法来计算最短路径
func findLadders(beginWord string, endWord string, wordList []string) []string {
	n := len(wordList)
	hasBegin, hasEnd := false, false
	for _, word := range wordList {
		if word == beginWord {
			hasBegin = true
		}
		if word == endWord {
			hasEnd = true
		}
		if hasBegin && hasEnd {
			break
		}
	}
	if !hasEnd {
		return nil
	}
	if !hasBegin {
		wordList = append(wordList, beginWord)
		n++
	}

	paths := make(map[string][]string, n)
	graph := make(map[string][]string, n)
	var q []string
	for i := 0; i < n; i++ {
		u := wordList[i]
		for j := i + 1; j < n; j++ {
			v := wordList[j]
			if isAdjacent(u, v) {
				graph[u] = append(graph[u], v)
				graph[v] = append(graph[v], u)
				if u == beginWord || v == beginWord {
					w := u
					if w == beginWord {
						w = v
					}
					paths[w] = []string{beginWord, w}
					q = append(q, w)
				}
			}
		}
	}
	for i := 0; i < len(q); i++ {
		u := q[i]
		for _, v := range graph[u] {
			if len(paths[v]) == 0 || len(paths[v]) > len(paths[u])+1 {
				q = append(q, v)
				paths[v] = make([]string, len(paths[u])+1)
				copy(paths[v], paths[u])
				paths[v][len(paths[u])] = v
			}
		}
	}
	return paths[endWord]
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
