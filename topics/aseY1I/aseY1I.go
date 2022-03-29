// 剑指 Offer II 005. 单词长度的最大乘积
//给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
//
//
//
//示例 1:
//
//输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
//输出: 16
//解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
//示例 2:
//
//输入: words = ["a","ab","abc","d","cd","bcd","abcd"]
//输出: 4
//解释: 这两个单词为 "ab", "cd"。
//示例 3:
//
//输入: words = ["a","aa","aaa","aaaa"]
//输出: 0
//解释: 不存在这样的两个单词。
//
//
//提示：
//
//2 <= words.length <= 1000
//1 <= words[i].length <= 1000
//words[i] 仅包含小写字母
//
//
//注意：本题与主站 318 题相同：https://leetcode-cn.com/problems/maximum-product-of-word-lengths/
//
//通过次数21,118提交次数29,954
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/29 下午5:06
package aseY1I

// 使用数组来记录每个单词包含的字母
func maxProduct(words []string) int {
	dict := map[int][26]bool{}
	for i, word := range words {
		chars := [26]bool{}
		for _, r := range word {
			chars[r-'a'] = true
		}
		dict[i] = chars
	}
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			hasSame := false
			for k := 0; k < 26; k++ {
				if dict[i][k] && dict[j][k] {
					hasSame = true
					break
				}
			}
			if !hasSame {
				product := len(words[i]) * len(words[j])
				if product > max {
					max = product
				}
			}
		}
	}
	return max
}

// 使用一个int的二进制位来记录每个单词是否包含各个字母
func maxProduct2(words []string) int {
	dict := map[int]int{}
	for i, word := range words {
		chars := 0
		for _, r := range word {
			chars |= 1 << (r - 'a')
		}
		dict[i] = chars
	}
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if dict[i]&dict[j] == 0 {
				product := len(words[i]) * len(words[j])
				if product > max {
					max = product
				}
			}
		}
	}
	return max
}

// 基于方法2的优化，只记录每种bitmask对应的最长单词
func maxProduct3(words []string) int {
	dict := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, r := range word {
			mask |= 1 << (r - 'a')
		}
		if len(word) > dict[mask] {
			dict[mask] = len(word)
		}
	}
	max := 0
	for mask1, len1 := range dict {
		for mask2, len2 := range dict {
			if mask1&mask2 == 0 && len1*len2 > max {
				max = len1 * len2
			}
		}
	}
	return max
}

// 基于方法3的优化，将map转为数组从而避免重复比较
func maxProduct4(words []string) int {
	dict := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, r := range word {
			mask |= 1 << (r - 'a')
		}
		if len(word) > dict[mask] {
			dict[mask] = len(word)
		}
	}
	type maskInfo struct {
		bit, len int
	}
	masks := make([]maskInfo, 0, len(dict))
	for mask, length := range dict {
		masks = append(masks, maskInfo{mask, length})
	}
	max := 0
	for i := 0; i < len(masks); i++ {
		for j := i + 1; j < len(masks); j++ {
			if masks[i].bit&masks[j].bit == 0 && masks[i].len*masks[j].len > max {
				max = masks[i].len * masks[j].len
			}
		}
	}
	return max
}
