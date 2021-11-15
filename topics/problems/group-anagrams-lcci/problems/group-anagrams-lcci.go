// 面试题 10.02. 变位词组
//编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
//
//注意：本题相对原题稍作修改
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。
//通过次数35,118提交次数47,447
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/15/21 11:00 AM
package problems

import "sort"

// 排序
func groupAnagrams(strs []string) [][]string {
	words := map[string][]string{}
	for _, str := range strs {
		key := reorder(str)
		words[key] = append(words[key], str)
	}
	res := make([][]string, 0, len(words))
	for _, vals := range words {
		res = append(res, vals)
	}
	return res
}
func reorder(s string) string {
	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

// 计数
func groupAnagrams2(strs []string) [][]string {
	words := map[[26]int][]string{}
	for _, str := range strs {
		key := getCount(str)
		words[key] = append(words[key], str)
	}
	res := make([][]string, 0, len(words))
	for _, vals := range words {
		res = append(res, vals)
	}
	return res
}

func getCount(str string) [26]int {
	count := [26]int{}
	for _, char := range str {
		count[char-'a']++
	}
	return count
}
