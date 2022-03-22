// 剑指 Offer II 015. 字符串中的所有变位词
//给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//变位词 指字母相同，但排列不同的字符串。
//
//
//
//示例 1:
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。
// 示例 2:
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的变位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的变位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的变位词。
//
//
//提示:
//
//1 <= s.length, p.length <= 3 * 104
//s 和 p 仅包含小写字母
//
//
//注意：本题与主站 438 题相同： https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
//
//通过次数11,642提交次数18,671
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/22 上午10:43
package VabMRr

func findAnagrams(s string, p string) []int {
	m, n := len(p), len(s)
	if m > n {
		return nil
	}
	var ans []int
	pChars := map[rune]int{}
	for _, r := range p {
		pChars[r]++
	}
	sChars := map[rune]int{}
	for _, r := range s[:m] {
		sChars[r]++
	}
	if isAnagram(pChars, sChars) {
		ans = append(ans, 0)
	}
	for i := 1; i+m <= len(s); i++ {
		sChars[rune(s[i-1])]--
		sChars[rune(s[i+m-1])]++
		if isAnagram(pChars, sChars) {
			ans = append(ans, i)
		}
	}
	return ans
}

func isAnagram(chars1, chars2 map[rune]int) bool {
	for r := range chars1 {
		if chars2[r] != chars1[r] {
			return false
		}
	}
	return true
}

// 参考了优秀题解
func findAnagrams2(s string, p string) []int {
	var ans []int
	if len(p) > len(s) {
		return ans
	}
	nums := [26]int{}
	for i := 0; i < len(p); i++ {
		nums[p[i]-'a']--
	}
	left := 0
	for right, r := range s {
		nums[r-'a']++
		for nums[r-'a'] > 0 {
			nums[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) {
			ans = append(ans, left)
		}
	}
	return ans
}
