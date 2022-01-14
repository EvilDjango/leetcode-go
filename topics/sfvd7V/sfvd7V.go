// 剑指 Offer II 033. 变位词组
//给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。
//
//注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。
//
//
//
//示例 1:
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//示例 2:
//
//输入: strs = [""]
//输出: [[""]]
//示例 3:
//
//输入: strs = ["a"]
//输出: [["a"]]
//
//
//提示：
//
//1 <= strs.length <= 104
//0 <= strs[i].length <= 100
//strs[i] 仅包含小写字母
//
//
//注意：本题与主站 49 题相同： https://leetcode-cn.com/problems/group-anagrams/
//
//通过次数5,915提交次数7,726
//
// @author xuejunc deerhunter0837@gmail.com
// @create 1/9/22 2:54 PM
package sfvd7V

import "sort"

// 会超时
func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	for _, str := range strs {
		i := 0
		for ; i < len(ans); i++ {
			if isAnagram(str, ans[i][0]) {
				ans[i] = append(ans[i], str)
				break
			}
		}
		if i == len(ans) {
			ans = append(ans, []string{str})
		}
	}
	return ans
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	count := map[rune]int{}
	for _, b := range a {
		count[b]++
	}
	for _, b := range b {
		count[b]--
		if count[b] == 0 {
			delete(count, b)
		}
	}
	return len(count) == 0
}

func groupAnagrams2(strs []string) [][]string {
	getKey := func(str string) string {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		return string(bytes)
	}
	anagrams := map[string][]string{}
	for _, str := range strs {
		key := getKey(str)
		anagrams[key] = append(anagrams[key], str)
	}
	ans := make([][]string, len(anagrams))
	i := 0
	for _, words := range anagrams {
		ans[i] = words
		i++
	}
	return ans
}
