package leetcode_go

/*
249. 移位字符串分组
给定一个字符串，对该字符串可以进行 “移位” 的操作，也就是将字符串中每个字母都变为其在字母表中后续的字母，比如："abc" -> "bcd"。这样，我们可以持续进行 “移位” 操作，从而生成如下移位序列：

"abc" -> "bcd" -> ... -> "xyz"
给定一个包含仅小写字母字符串的列表，将该列表中所有满足 “移位” 操作规律的组合进行分组并返回。



示例：

输入：["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"]
输出：
[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]
解释：可以认为字母表首尾相接，所以 'z' 的后续为 'a'，所以 ["az","ba"] 也满足 “移位” 操作规律。
通过次数5,117提交次数8,061

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/22/21 7:51 PM
*/

func groupStrings(strings []string) [][]string {
	var ans [][]string
	grouped := make([]bool, len(strings))
	for i := 0; i < len(strings); i++ {
		if grouped[i] {
			continue
		}
		arr := []string{strings[i]}
		for j := i + 1; j < len(strings); j++ {
			if grouped[j] {
				continue
			}
			if canTransform(strings[i], strings[j]) {
				arr = append(arr, strings[j])
				grouped[j] = true
			}
		}
		ans = append(ans, arr)
	}
	return ans
}

func canTransform(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 1; i < len(s); i++ {
		if (int(s[i])-int(s[i-1])-int(t[i])+int(t[i-1]))%26 != 0 {
			return false
		}
	}
	return true
}
