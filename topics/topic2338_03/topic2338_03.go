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

//提交记录中看到的优秀解法
func findLadders(beginWord string, endWord string, wordList []string) []string {
	n := len(wordList)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = true
	}
	for i := range wordList {
		if wordList[i] == beginWord {
			vis[i] = false
		}
	}
	var ans []string
	var dfs func(str string) bool
	dfs = func(str string) bool {
		if str == endWord {
			ans = append([]string{str}, ans...)
			return true
		}
		for i := 0; i < n; i++ {
			if vis[i] && judge(wordList[i], str) {
				vis[i] = false
				if dfs(wordList[i]) {
					ans = append([]string{str}, ans...)
					return true
				}
			}
		}
		return false
	}
	dfs(beginWord)
	return ans
}

func judge(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
