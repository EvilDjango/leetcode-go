// 面试题 17.13. 恢复空格
//哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。
//
//注意：本题相对原题稍作改动，只需返回未识别的字符数
//
//
//
//示例：
//
//输入：
//dictionary = ["looked","just","like","her","brother"]
//sentence = "jesslookedjustliketimherbrother"
//输出： 7
//解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
//提示：
//
//0 <= len(sentence) <= 1000
//dictionary中总字符数不超过 150000。
//你可以认为dictionary和sentence中只包含小写字母。
//通过次数24,412提交次数44,153
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/17/21 10:25 AM
package re_space_lcci

// 暴力回溯，会超时
func respace(dictionary []string, sentence string) int {
	ans := len(sentence)
	dfs(dictionary, sentence, 0, 0, &ans)
	return ans
}

func dfs(dictionary []string, sentence string, index, unknown int, minUnknown *int) {
	if unknown >= *minUnknown {
		return
	}
	if index == len(sentence) {
		*minUnknown = unknown
		return
	}
	for _, w := range dictionary {
		if index+len(w) > len(sentence) || sentence[index:index+len(w)] != w {
			continue
		}
		dfs(dictionary, sentence, index+len(w), unknown, minUnknown)
	}
	dfs(dictionary, sentence, index+1, unknown+1, minUnknown)
}

// 动态规划
func respace2(dictionary []string, sentence string) int {
	length := len(sentence)
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		minUnknown := dp[i-1] + 1
		for _, w := range dictionary {
			if i < len(w) || sentence[i-len(w):i] != w {
				continue
			}
			if dp[i-len(w)] < minUnknown {
				minUnknown = dp[i-len(w)]
			}
		}
		dp[i] = minUnknown
	}
	return dp[length]
}

type Trie struct {
	IsLeaf bool
	Next   [26]*Trie
}

func (t *Trie) Insert(s string) {
	curr := t
	// 因为是后缀树，所以这里是把单词倒序插入字典树的
	for i := len(s) - 1; i >= 0; i-- {
		index := s[i] - 'a'
		if curr.Next[index] == nil {
			curr.Next[index] = &Trie{}
		}
		curr = curr.Next[index]
	}
	curr.IsLeaf = true
}

// 后缀树+动态规划
func respace3(dictionary []string, sentence string) int {
	root := &Trie{}
	for _, w := range dictionary {
		root.Insert(w)
	}
	length := len(sentence)
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		dp[i] = dp[i-1] + 1
		curr := root
		for j := i - 1; j >= 0; j-- {
			curr = curr.Next[sentence[j]-'a']
			if curr == nil {
				break
			}
			if curr.IsLeaf {
				if dp[j] < dp[i] {
					dp[i] = dp[j]
				}
			}
		}
	}
	return dp[length]
}
