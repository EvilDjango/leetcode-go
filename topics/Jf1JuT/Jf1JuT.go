// 剑指 Offer II 114. 外星文字典
//现有一种使用英语字母的外星文语言，这门语言的字母顺序与英语顺序不同。
//
//给定一个字符串列表 words ，作为这门语言的词典，words 中的字符串已经 按这门新语言的字母顺序进行了排序 。
//
//请你根据该词典还原出此语言中已知的字母顺序，并 按字母递增顺序 排列。若不存在合法字母顺序，返回 "" 。若存在多种可能的合法字母顺序，返回其中 任意一种 顺序即可。
//
//字符串 s 字典顺序小于 字符串 t 有两种情况：
//
//在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于 t 中字母之前，那么 s 的字典顺序小于 t 。
//如果前面 min(s.length, t.length) 字母都相同，那么 s.length < t.length 时，s 的字典顺序也小于 t 。
//
//
//示例 1：
//
//输入：words = ["wrt","wrf","er","ett","rftt"]
//输出："wertf"
//示例 2：
//
//输入：words = ["z","x"]
//输出："zx"
//示例 3：
//
//输入：words = ["z","x","z"]
//输出：""
//解释：不存在合法字母顺序，因此返回 "" 。
//
//
//提示：
//
//1 <= words.length <= 100
//1 <= words[i].length <= 100
//words[i] 仅由小写英文字母组成
//
//
//注意：本题与主站 269 题相同： https://leetcode-cn.com/problems/alien-dictionary/
//
//通过次数1,026提交次数2,416
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/11/21 12:31 PM
package Jf1JuT

func alienOrder(words []string) string {
	edges := map[byte][]byte{}
	inDegree := map[byte]int{}
	seen := map[byte]bool{}
	count := 0
	for i := 0; i < len(words); i++ {
		w1, w2 := "", words[i]
		if i > 0 {
			w1 = words[i-1]
		}
		foundDiff := false
		for j := 0; j < len(w1) || j < len(w2); j++ {
			u, v := byte(' '), byte(' ')
			if j < len(w1) {
				u = w1[j]
				if !seen[u] {
					count++
					seen[u] = true
					inDegree[u] = 0
				}
			}
			if j < len(w2) {
				v = w2[j]
				if !seen[v] {
					count++
					seen[v] = true
					inDegree[v] = 0
				}
			}
			if !foundDiff && j < len(w1) && j < len(w2) && u != v {
				edges[u] = append(edges[u], v)
				inDegree[v]++
				foundDiff = true
			}
		}
		// 非法
		if !foundDiff && len(w1) > len(w2) {
			return ""
		}
	}
	var zeroDegree []byte
	for b, d := range inDegree {
		if d == 0 {
			zeroDegree = append(zeroDegree, b)
		}
	}
	var bytes []byte
	for len(zeroDegree) > 0 {
		b := zeroDegree[0]
		zeroDegree = zeroDegree[1:]
		bytes = append(bytes, b)
		for _, to := range edges[b] {
			inDegree[to]--
			if inDegree[to] == 0 {
				zeroDegree = append(zeroDegree, to)
			}
		}
	}
	if len(bytes) < count {
		return ""
	}
	return string(bytes)
}
