package leetcode_go

/*
269. 火星词典
现有一种使用英语字母的火星语言，这门语言的字母顺序与英语顺序不同。

给你一个字符串列表 words ，作为这门语言的词典，words 中的字符串已经 按这门新语言的字母顺序进行了排序 。

请你根据该词典还原出此语言中已知的字母顺序，并 按字母递增顺序 排列。若不存在合法字母顺序，返回 "" 。若存在多种可能的合法字母顺序，返回其中 任意一种 顺序即可。

字符串 s 字典顺序小于 字符串 t 有两种情况：

在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于 t 中字母之前，那么 s 的字典顺序小于 t 。
如果前面 min(s.length, t.length) 字母都相同，那么 s.length < t.length 时，s 的字典顺序也小于 t 。


示例 1：

输入：words = ["wrt","wrf","er","ett","rftt"]
输出："wertf"
示例 2：

输入：words = ["z","x"]
输出："zx"
示例 3：

输入：words = ["z","x","z"]
输出：""
解释：不存在合法字母顺序，因此返回 "" 。


提示：

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] 仅由小写英文字母组成
通过次数5,723提交次数16,811
Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/2/21 11:34 AM
*/

func alienOrder(words []string) string {
	n := len(words)

	var edges [][]int
	for i := 1; i < n; i++ {
		u, v, ok := getEdge(words[i-1], words[i])
		if ok {
			edges = append(edges, []int{u, v})
		} else if len(words[i-1]) > len(words[i]) {
			// 当两个词前面部分完全相同，且第一个词的长度比较小，则输入排序有误
			return ""
		}
	}

	nodes := make([]bool, 26)
	for _, word := range words {
		for _, r := range word {
			nodes[r-'a'] = true
		}
	}

	topo := getTopo(edges, nodes)

	runes := make([]rune, len(topo))
	for i := 0; i < len(topo); i++ {
		runes[i] = rune('a' + topo[i])
	}
	return string(runes)
}

func getTopo(edges [][]int, nodes []bool) []int {
	adjacents := make([][]int, 26)
	inDeg := make([]int, 26)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjacents[u] = append(adjacents[u], v)
		inDeg[v]++
	}
	var q []int
	for i := 0; i < 26; i++ {
		// 未出现的字母忽略掉
		if inDeg[i] == 0 && nodes[i] {
			q = append(q, i)
		}
	}
	visitedEdges := 0
	for i := 0; i < len(q); i++ {
		for _, adj := range adjacents[q[i]] {
			inDeg[adj]--
			visitedEdges++
			if inDeg[adj] == 0 {
				q = append(q, adj)
			}
		}
	}
	if visitedEdges == len(edges) {
		return q
	}
	return nil
}

func getEdge(s1, s2 string) (int, int, bool) {
	minLen := Min(len(s1), len(s2))
	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] {
			return int(s1[i] - 'a'), int(s2[i] - 'a'), true
		}
	}
	return 0, 0, false
}

// 简洁版
func alienOrder2(words []string) string {
	adjacency := make(map[uint8][]uint8)
	inDegrees := make(map[uint8]int)

	for _, word := range words {
		for i := range word {
			adjacency[word[i]] = []uint8{}
		}
	}
	for i := 1; i < len(words); i++ {
		w1, w2 := words[i-1], words[i]
		j := 0
		minLen := Min(len(w1), len(w2))
		for ; j < minLen; j++ {
			if w1[j] != w2[j] {
				adjacency[w1[j]] = append(adjacency[w1[j]], w2[j])
				inDegrees[w2[j]]++
				break
			}
		}
		if j == minLen && len(w1) > len(w2) {
			return ""
		}
	}

	var q []uint8
	for u := range adjacency {
		if inDegrees[u] == 0 {
			q = append(q, u)
		}
	}

	visited := 0
	for i := 0; i < len(q); i++ {
		cur := q[i]
		for _, adj := range adjacency[cur] {
			inDegrees[adj]--
			if inDegrees[adj] == 0 {
				q = append(q, adj)
			}
		}
		visited++
	}

	// 有环，不存在合法顺序
	if visited != len(adjacency) {
		return ""
	}
	return string(q)
}
