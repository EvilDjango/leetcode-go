// 面试题 17.15. 最长单词
//给定一组单词words，编写一个程序，找出其中的最长单词，且该单词由这组单词中的其他单词组合而成。若有多个长度相同的结果，返回其中字典序最小的一项，若没有符合要求的单词则返回空字符串。
//
//示例：
//
//输入： ["cat","banana","dog","nana","walk","walker","dogwalker"]
//输出： "dogwalker"
//解释： "dogwalker"可由"dog"和"walker"组成。
//提示：
//
//0 <= len(words) <= 200
//1 <= len(words[i]) <= 100
//通过次数5,668提交次数14,077
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/16/21 4:12 PM
package topic2342

func longestWord(words []string) string {
	ans := ""
	for _, word := range words {
		if (len(word) > len(ans) || len(word) == len(ans) && word < ans) && isCombined(word, words, 0) {
			ans = word
		}
	}
	return ans
}

func isCombined(word string, words []string, length int) bool {
	if length == len(word) {
		return true
	}

	for _, w := range words {
		if w == word || length+len(w) > len(word) {
			continue
		}
		if word[length:length+len(w)] == w && isCombined(word, words, length+len(w)) {
			return true
		}
	}
	return false
}
