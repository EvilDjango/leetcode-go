// 面试题 01.02. 判定是否互为字符重排
//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
//示例 1：
//
//输入: s1 = "abc", s2 = "bca"
//输出: true
//示例 2：
//
//输入: s1 = "abc", s2 = "bad"
//输出: false
//说明：
//
//0 <= len(s1) <= 100
//0 <= len(s2) <= 100
//通过次数62,384提交次数96,982
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 11:23 PM
package check_permutation_lcci

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := map[rune]int{}
	for _, r := range s1 {
		count[r]++
	}
	for _, r := range s2 {
		count[r]--
		if count[r] == 0 {
			delete(count, r)
		}
	}
	return len(count) == 0
}
