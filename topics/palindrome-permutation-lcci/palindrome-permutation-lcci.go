// 面试题 01.04. 回文排列
//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
//回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//
//回文串不一定是字典当中的单词。
//
//
//
//示例1：
//
//输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
//
//
//通过次数43,309提交次数78,872
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 10:37 PM
package palindrome_permutation_lcci

func canPermutePalindrome(s string) bool {
	seen := map[rune]bool{}
	for _, r := range s {
		if seen[r] {
			delete(seen, r)
		} else {
			seen[r] = true
		}
	}
	return len(seen) < 2

}
