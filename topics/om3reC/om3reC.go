// 剑指 Offer II 108. 单词演变
//在字典（单词列表） wordList 中，从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
//
//序列中第一个单词是 beginWord 。
//序列中最后一个单词是 endWord 。
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典 wordList 中的单词。
//给定两个长度相同但内容不同的单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。
//
//
//
//示例 1：
//
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
//输出：5
//解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
//示例 2：
//
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
//输出：0
//解释：endWord "cog" 不在字典中，所以无法进行转换。
//
//
//提示：
//
//1 <= beginWord.length <= 10
//endWord.length == beginWord.length
//1 <= wordList.length <= 5000
//wordList[i].length == beginWord.length
//beginWord、endWord 和 wordList[i] 由小写英文字母组成
//beginWord != endWord
//wordList 中的所有字符串 互不相同
//
//
//注意：本题与主站 127 题相同： https://leetcode-cn.com/problems/word-ladder/
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/15/21 6:30 PM
package om3reC

// 双向广度遍历
func ladderLength(beginWord string, endWord string, wordList []string) int {
	foundEnd := false
	for _, word := range wordList {
		if word == endWord {
			foundEnd = true
		}
		if foundEnd {
			break
		}
	}
	if !foundEnd {
		return 0
	}
	forward := []string{beginWord}
	backward := []string{endWord}
	fVisited := map[string]bool{beginWord: true}
	bVisited := map[string]bool{endWord: true}
	steps := 1
	for len(forward) > 0 {
		if len(forward) > len(backward) {
			forward, backward = backward, forward
			fVisited, bVisited = bVisited, fVisited
		}
		steps++
		size := len(forward)
		for i := 0; i < size; i++ {
			for _, word := range wordList {
				if fVisited[word] {
					continue
				}
				if !canTransform(forward[i], word) {
					continue
				}
				if bVisited[word] {
					return steps
				}
				forward = append(forward, word)
				fVisited[word] = true
			}
		}
		forward = forward[size:]
	}
	return 0
}

func canTransform(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return true
}

// 双向广度遍历+优化建图
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	ids := map[string]int{}
	var graph [][]int
	addWord := func(s string) int {
		id, ok := ids[s]
		if !ok {
			id = len(graph)
			ids[s] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(s string) {
		id := addWord(s)
		bytes := []byte(s)
		for i, b := range bytes {
			bytes[i] = '*'
			w := string(bytes)
			id2 := addWord(w)
			graph[id] = append(graph[id], id2)
			graph[id2] = append(graph[id2], id)
			bytes[i] = b
		}
	}
	for _, word := range wordList {
		addEdge(word)
	}
	addEdge(beginWord)
	if _, ok := ids[endWord]; !ok {
		return 0
	}
	forward := []int{ids[beginWord]}
	backward := []int{ids[endWord]}
	disForward := make([]int, len(ids))
	for i := 0; i < len(ids); i++ {
		disForward[i] = -1
	}
	disBackward := make([]int, len(ids))
	copy(disBackward, disForward)
	disForward[ids[beginWord]] = 0
	disBackward[ids[endWord]] = 0
	for len(forward) > 0 {
		if len(forward) > len(backward) {
			forward, backward = backward, forward
			disForward, disBackward = disBackward, disForward
		}
		size := len(forward)
		for i := 0; i < size; i++ {
			id := forward[i]
			for _, next := range graph[id] {
				if disForward[next] != -1 {
					continue
				}
				if disBackward[next] != -1 {
					return (disForward[id] + disBackward[next] + 3) / 2
				}
				disForward[next] = disForward[id] + 1
				forward = append(forward, next)
			}
		}
		forward = forward[size:]
	}
	return 0
}
